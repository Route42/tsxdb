package backend

import (
	"github.com/Route42/tsxdb/rpc/types"
)

type IMetadata interface {
	// @todo params, responses etc
	CreateOrUpdateSeries(*CreateSeries) *CreateSeriesResult // create/update new series (batch)
	SearchSeries(*SearchSeries) *SearchSeriesResult         // search one or multiple series by tags
	DeleteSeries(*DeleteSeries) *DeleteSeriesResult         // remove series (batch)
}

type CreateSeries struct {
	Series map[types.SeriesCreateIdentifier]types.SeriesCreateMetadata
}

type CreateSeriesResult struct {
	Results map[types.SeriesCreateIdentifier]types.SeriesMetadataResponse
	Error   error
}

type SearchSeries struct {
	SearchSeriesElement
}

type SearchSeriesResult struct {
	Series []types.SeriesIdentifier
	Error  error
}

type DeleteSeries struct {
	Series []types.SeriesIdentifier
}

type DeleteSeriesResult struct {
	Error error
}

type SearchSeriesElement struct {
	Namespace  int
	Name       string
	Tag        string
	Comparator SearchSeriesComparator
	And        []SearchSeriesElement
	Or         []SearchSeriesElement
}

type SearchSeriesComparator string

const SearchSeriesComparatorEquals SearchSeriesComparator = "EQUALS"
const SearchSeriesComparatorNot SearchSeriesComparator = "NOT"
