//go:generate go run ../../generate/tags/main.go -ListTagsInIDElem=PipelineId -ServiceTagsSlice -TagOp=AddTags -TagInIDElem=PipelineId -UntagOp=RemoveTags -UpdateTags
//go:generate go run ../../generate/servicepackage/main.go
// ONLY generate directives and package declaration! Do not add anything else to this file.

package datapipeline
