package service

import (
	"CN-EU-FSIMS/internal/app/models"
	"CN-EU-FSIMS/internal/app/models/query"
	"context"
	"errors"

	"strings"

	"github.com/golang/glog"
)

func PushNotification(receiver string, event *Event) error {
	err := CheckUserIsExisted(receiver)
	if err != nil {
		return err
	}

	notification := models.Notification{
		Source:    event.Source,
		Content:   event.Content,
		UUID:      receiver,
		EventTime: event.EventTime,
		EventType: event.EventType,
		Affected:  event.AffectedBatchNumber,
		Proposal:  event.Proposal,
		RiskLevel: event.RiskLevel,
		State:     1,
	}

	err = query.Q.Notification.WithContext(context.Background()).Create(&notification)
	if err != nil {
		return err
	}

	return nil
}

func ReadNotification(id int) error {
	q := query.Q.Notification

	_, err := q.WithContext(context.Background()).Where(q.ID.Eq(uint(id))).Update(q.State, 2)
	if err != nil {
		return err
	}
	return nil
}

func GetNotificationCount(uuid string) (int, error) {
	var count = 0
	q := query.Q.Notification
	res, err := q.WithContext(context.Background()).Where(q.UUID.Eq(uuid)).Find()
	if err != nil {
		return 0, err
	}
	for _, v := range res {
		if v.State == 1 {
			count++
		}
	}
	return count, nil
}

func GetNotification(uuid string) ([]models.ResNotification, int, error) {

	q := query.Q.Notification
	res, err := q.WithContext(context.Background()).Where(q.UUID.Eq(uuid)).Find()
	if err != nil {
		return nil, 0, err
	}

	count := len(res)
	records := make([]models.ResNotification, count)
	for i, v := range res {
		glog.Infoln("database：", res[i])
		var data = models.ToNotificationInfo(v)
		//var house_number
		res, _ := ToResNotification(data)
		// glog.Infoln("test:", res)
		records[i] = *res
	}

	return records, count, nil

}

func ToResNotification(data models.NotificationInfo) (*models.ResNotification, error) {
	house_number := data.Source
	//affected := data.Affected
	//regex := regexp.MustCompile(`\[BATCH-[^\]]+\]`)

	// parts := house_number.split('-')
	parts := strings.Split(house_number, "-")
	prefix := parts[0]
	var res *models.ResNotification
	glog.Infoln("prefix:", prefix)
	switch prefix {
	case "PASP":
		res, err := PastureNotification(data)
		if err != nil {
			return nil, err
		}
		glog.Infoln("res:", res)
		return res, nil
		// return res,nil
	case "SLAHP":
		res, err := SlaughterNotification(data)
		if err != nil {
			return nil, err
		}
		glog.Infoln("res:", res)
		return res, nil
	case "PACHP":
		res, err := PackageNotification(data)
		if err != nil {
			return nil, err
		}
		glog.Infoln("res:", res)
		return res, nil
	case "VEH":
		res, err := VechicalNotification(data)
		if err != nil {
			return nil, err
		}
		glog.Infoln("res:", res)
		return res, nil
	default:
		glog.Infoln("Unknown prefix")

	}

	return res, errors.New("通知解析失败")
}

func PastureNotification(data models.NotificationInfo) (*models.ResNotification, error) {
	house_number := data.Source
	affected := data.Affected
	//regex := regexp.MustCompile(`\[BATCH-[^\]]+\]`)
	res, err := query.PastureHouse.WithContext(context.Background()).Where(query.PastureHouse.HouseNumber.Eq(house_number)).First()
	if err != nil {
		return nil, err
	}
	source_name := res.Name
	affected = affected[1 : len(affected)-2]
	matches := strings.Split(affected, "];[")
	//matches := regex.FindAllString(affected, -1)
	glog.Infoln("matches:", matches)
	batches := make([]models.Batch, 0)
	for _, v := range matches {
		batch_number := v
		batch, err := query.FeedingBatch.WithContext(context.Background()).Where(query.FeedingBatch.BatchNumber.Eq(batch_number)).First()
		if err != nil {
			return nil, err
		}
		//state := batch.State
		batch_res := models.Batch{
			BatchNumber: batch_number,
			State:       batch.State,
		}
		batches = append(batches, batch_res)
	}
	glog.Infoln("batches", batches)
	not := models.ResNotification{
		ID:         data.ID,
		SourceName: source_name,
		Content:    data.Content,
		UUID:       data.UUID,
		NoticeTime: data.NoticeTime,
		EventTime:  data.EventTime,
		Affected:   batches,
		Proposal:   data.Proposal,
		RiskLevel:  data.RiskLevel,
		State:      data.State,
	}
	return &not, nil
}

func SlaughterNotification(data models.NotificationInfo) (*models.ResNotification, error) {
	house_number := data.Source
	affected := data.Affected
	//regex := regexp.MustCompile(`\[BATCH-[^\]]+\]`)
	res, err := query.SlaughterHouse.WithContext(context.Background()).Where(query.SlaughterHouse.HouseNumber.Eq(house_number)).First()
	if err != nil {
		return nil, err
	}
	source_name := res.Name
	//affected = affected[1 : len(affected)-2]
	matches := strings.Split(affected, ";")

	//matches := regex.FindAllString(affected, -1)
	glog.Infoln("matches:", matches)
	batches := make([]models.Batch, 0)
	for _, v := range matches[:len(matches)-1] {
		batch_number := v[1 : len(v)-1]
		glog.Infoln("batch_number:", batch_number)
		batch, err := query.SlaughterBatch.WithContext(context.Background()).Where(query.SlaughterBatch.BatchNumber.Eq(batch_number)).First()
		if err != nil {
			return nil, err
		}
		//state := batch.State
		batch_res := models.Batch{
			BatchNumber: batch_number,
			State:       batch.State,
		}
		batches = append(batches, batch_res)
	}
	glog.Infoln("batches", batches)
	not := models.ResNotification{
		ID:         data.ID,
		SourceName: source_name,
		Content:    data.Content,
		UUID:       data.UUID,
		NoticeTime: data.NoticeTime,
		EventTime:  data.EventTime,
		Affected:   batches,
		Proposal:   data.Proposal,
		RiskLevel:  data.RiskLevel,
		State:      data.State,
	}
	return &not, nil
}

func PackageNotification(data models.NotificationInfo) (*models.ResNotification, error) {
	house_number := data.Source
	affected := data.Affected
	//regex := regexp.MustCompile(`\[BATCH-[^\]]+\]`)
	res, err := query.PackageHouse.WithContext(context.Background()).Where(query.PackageHouse.HouseNumber.Eq(house_number)).First()
	if err != nil {
		return nil, err
	}
	source_name := res.Name
	affected = affected[1 : len(affected)-2]
	matches := strings.Split(affected, "];[")
	//matches := regex.FindAllString(affected, -1)
	glog.Infoln("matches:", matches)
	batches := make([]models.Batch, 0)
	for _, v := range matches {
		batch_number := v
		batch, err := query.PackageBatch.WithContext(context.Background()).Where(query.PackageBatch.BatchNumber.Eq(batch_number)).First()
		if err != nil {
			return nil, err
		}
		//state := batch.State
		batch_res := models.Batch{
			BatchNumber: batch_number,
			State:       batch.State,
		}
		batches = append(batches, batch_res)
	}
	glog.Infoln("batches", batches)
	not := models.ResNotification{
		ID:         data.ID,
		SourceName: source_name,
		Content:    data.Content,
		UUID:       data.UUID,
		NoticeTime: data.NoticeTime,
		EventTime:  data.EventTime,
		Affected:   batches,
		Proposal:   data.Proposal,
		RiskLevel:  data.RiskLevel,
		State:      data.State,
	}
	return &not, nil
}

func VechicalNotification(data models.NotificationInfo) (*models.ResNotification, error) {
	tv_number := data.Source
	affected := data.Affected
	//regex := regexp.MustCompile(`\[BATCH-[^\]]+\]`)
	res, err := query.TransportVehicle.WithContext(context.Background()).Where(query.TransportVehicle.TVNumber.Eq(tv_number)).First()
	if err != nil {
		return nil, err
	}
	source_name := res.LicenseNumber
	affected = affected[1 : len(affected)-2]
	matches := strings.Split(affected, "];[")
	//matches := regex.FindAllString(affected, -1)
	glog.Infoln("matches:", matches)
	batches := make([]models.Batch, 0)
	for _, v := range matches {
		batch_number := v
		batch, err := query.TransportBatch.WithContext(context.Background()).Where(query.TransportBatch.BatchNumber.Eq(batch_number)).First()
		if err != nil {
			return nil, err
		}
		//state := batch.State
		batch_res := models.Batch{
			BatchNumber: batch_number,
			State:       batch.State,
		}
		batches = append(batches, batch_res)
	}
	glog.Infoln("batches", batches)
	not := models.ResNotification{
		ID:         data.ID,
		SourceName: source_name,
		Content:    data.Content,
		UUID:       data.UUID,
		NoticeTime: data.NoticeTime,
		EventTime:  data.EventTime,
		Affected:   batches,
		Proposal:   data.Proposal,
		RiskLevel:  data.RiskLevel,
		State:      data.State,
	}
	return &not, nil
}
