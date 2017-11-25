package chessclock

import "github.com/OpenPeeDeeP/chessclock/chessclock"

//NilStore does absolutly nothing to store the logs
type NilStore struct{}

//Start does absolutly nothing
func (s *NilStore) Start(timestamp int64, tag, description string) error {
	return nil
}

//Stop does absolutly nothing
func (s *NilStore) Stop(reason chessclock.StopRequest_Reason) error {
	return nil
}

//TimeSheets does absolutly nothing
func (s *NilStore) TimeSheets() ([]int64, error) {
	return nil, nil
}

//Events does absolutly nothing
func (s *NilStore) Events(date int64) ([]*Event, error) {
	return nil, nil
}
