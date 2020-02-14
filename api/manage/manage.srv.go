package manage

import "capulus/database/models"

type Host struct {
	ID   int
	Host string
	Time string
}

func (h *Host) GetAll() ([]models.Host, error) {
	hosts, err := models.GetHosts()
	if err != nil {
		return nil, err
	}

	return hosts, nil
}

func (h *Host) Add() error {
	return models.AddHost(
		h.Host,
		h.Time,
	)
}

func (h *Host) Delete() error {
	return models.DeleteHost(
		h.ID,
	)
}
