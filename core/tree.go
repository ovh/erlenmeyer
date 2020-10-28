package core

import (
	"fmt"

	"github.com/prometheus/common/model"
)

type payload interface{}

// Node is the struct holding the query tree.
// Needed for promql, it served also for Graphite protocol
type Node struct {
	Level   int
	Left    *Node
	Right   *Node
	Payload payload
}

// NewNode is creating a new node with a specific type and Payload
func NewNode(payload payload) *Node {
	return &Node{
		Left:    nil,
		Right:   nil,
		Payload: payload,
	}
}

// PrintNode is printing a tree
func (n *Node) PrintNode(level int) {
	if n == nil {
		return
	}
	fmt.Printf("%T\t", n.Payload)
	fmt.Printf("%+v\n", n.Payload)
	if n.Left != nil {
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Print("-- left:")
		n.Left.PrintNode(level + 1)
	}
	if n.Right != nil {
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Print("-- rigth:")
		n.Right.PrintNode(level + 1)
	}
}

// NewEmptyNode is a constructor
func NewEmptyNode() *Node {
	return NewNode(WarpScriptPayload{
		WarpScript: "",
	})
}

// FetchPayload is the payload for the fetch function
type FetchPayload struct {
	ClassName string
	Labels    map[string]string
	End       string
	Start     string
	Step      string
	Offset    string
	Absent    bool
	Instant   bool
}

// AggregatePayload represents an aggregation operation on a vector.
type AggregatePayload struct {
	Op               string   // The used aggregation operation.
	Param            string   // Parameter used by some aggregators.
	Grouping         []string // The labels by which to group the vector.
	Without          bool     // Whether to drop the given labels rather than keep them.
	KeepCommonLabels bool     // Whether to keep common labels among result elements.
}

// NumberLiteralPayload is holding a number
type NumberLiteralPayload struct {
	Value string
}

// BinaryExprPayload is hodling an Op
type BinaryExprPayload struct {
	Op             string
	IsOn           bool
	IsIgnoring     bool
	FilteredLabels []string
	IncludeLabels  []string
	Card           string
	ReturnBool     bool
}

// FunctionPayload represents a function of the expression language and is
// used by function nodes.
type FunctionPayload struct {
	Name         string
	ArgTypes     []model.ValueType
	Args         []string
	OptionalArgs int
}

// Context is holding the informations like token, start, end, and so on
type Context struct {
	Query string
	Start Time
	End   Time
	Step  string
}

// BucketizePayload is the payload for the bucketize function
type BucketizePayload struct {
	LastBucket   string
	BucketSpan   string
	BucketCount  string
	BucketRange  string
	PreBucketize string
	Filler       string
	Op           string
	Step         string
	ApplyRate    bool
	Absent       bool
}

// FindPayload is the payload for the find function
type FindPayload struct {
	ClassName string
	Labels    map[string]string
}

// WarpScriptPayload is the payload to push WarpScript directly in the tree
type WarpScriptPayload struct {
	WarpScript string
}

// StorePayload is the payload to store variable in the tree
type StorePayload struct {
	Name  string
	Value string
}

// MapperPayload is the payload to map GTS' value in the tree
type MapperPayload struct {
	Mapper      string
	Constant    string
	PreWindow   string
	PostWindow  string
	Occurrences string
}

// ReducerPayload is the payload to reduce GTS' value in the tree
type ReducerPayload struct {
	Reducer string
	Value   string
	Labels  []string
}

// AddValuePayload is the payload to add a value to a GTS
type AddValuePayload struct {
	Timestamp string
	Latitude  string
	Longitude string
	Elevation string
	Value     string
}

// FillValuePayload is the payload to fill empty value of a GTS
type FillValuePayload struct {
	Latitude  string
	Longitude string
	Elevation string
	Value     string
}
