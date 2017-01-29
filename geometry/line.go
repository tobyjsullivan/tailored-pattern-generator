package geometry

type Line interface {
    GetStart() *Point
    GetEnd() *Point

    ToEnglish() string
    ToAutoCAD() string
}

