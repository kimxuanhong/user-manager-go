package di

import (
	"fmt"
	"github.com/facebookgo/inject"
	"sync"
)

type DI interface {
	Dependencies(dependence interface{})
	Graph() error
}

type di struct {
	graph        inject.Graph
	dependencies []interface{}
}

var instanceDI *di
var dIOnce sync.Once

func NewDI() DI {
	dIOnce.Do(func() {
		instanceDI = &di{
			graph: inject.Graph{},
		}
	})
	return instanceDI
}

func (r *di) Dependencies(dependence interface{}) {
	r.dependencies = append(r.dependencies, dependence)
}

func (r *di) Graph() error {
	var objects []*inject.Object
	for _, dependence := range r.dependencies {
		objects = append(objects, &inject.Object{Value: dependence})
	}

	if err := r.graph.Provide(objects...); err != nil {
		fmt.Println("Error injecting:", err)
		return err
	}
	// Tạo các liên kết dependencies
	if err := r.graph.Populate(); err != nil {
		fmt.Println("Error populating:", err)
		return err
	}
	return nil
}
