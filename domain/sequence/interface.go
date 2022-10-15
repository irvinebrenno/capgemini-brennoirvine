package sequence

import "capgemini-brennoirvine/infrastructure/persistence/sequence"

// ISequence define uma interface para Sequence
type ISequence interface {
	CheckSequence(req *sequence.Sequence) (exists *bool, err error)
	InsertSequence(req *sequence.Sequence) (err error)
	GetStats() (res *sequence.Stats, err error)
	ConvertSequenceToData(latters []string) *sequence.Sequence
}
