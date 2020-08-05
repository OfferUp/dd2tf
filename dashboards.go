//go:generate go-bindata -o tpl.go tmpl

package main

import (
	"github.com/clarketm/go-datadog-api"
)

type Dashboard struct {
}

func (d Dashboard) getElement(client datadog.Client, id interface{}) (interface{}, error) {
	dash, err := client.GetDashboard(id)
	return dash, err
}

func (d Dashboard) getAsset() string {
	return "tmpl/timeboard.tmpl"
}

func (d Dashboard) getName() string {
	return "dashboards"
}

func (d Dashboard) String() string {
	return d.getName()
}

func (d Dashboard) getAllElements(client datadog.Client) ([]Item, error) {
	var ids []Item
	dashboards, err := client.GetDashboards()
	if err != nil {
		return ids, err
	}
	for _, elem := range dashboards {
		ids = append(ids, Item{id: *elem.Id, d: Dashboard{}})
	}
	return ids, nil
}
