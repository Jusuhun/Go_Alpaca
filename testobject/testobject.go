package testobject

type MotionProperty struct {
	velocity  float64
	acceltion float64
	juke      float64
	position  interface{}
}

type Motion struct {
	Kind     string
	Name     string
	Property MotionProperty
}

type Cylinder struct {
	Kind     string
	Name     string
	Property interface{}
}

type DataProperty struct {
	Kind     string
	Name     string
	Property interface{}
}

type Data struct {
	Kind     string
	Name     string
	Property DataProperty
}
