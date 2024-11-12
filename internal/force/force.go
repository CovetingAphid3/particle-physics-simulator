package force

type Force struct{
    Value float64
    XDirection float64
    YDirection float64
}

func NewForce(Value,xDirection,yDirection float64)*Force{
    return &Force{
        Value:Value,
        XDirection:xDirection,
        YDirection:yDirection,
    }
}
