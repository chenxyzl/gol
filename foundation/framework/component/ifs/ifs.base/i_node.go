package ifs_base

import (
	"foundation/framework/component"
)

type IComponentNode interface {
	GetComponent(comType component.ComType) IComponent
}
