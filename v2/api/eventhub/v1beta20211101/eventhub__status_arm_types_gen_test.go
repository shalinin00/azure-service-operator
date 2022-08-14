// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Eventhub_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Eventhub_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventhubStatusARM, EventhubStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventhubStatusARM runs a test to see if a specific instance of Eventhub_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventhubStatusARM(subject Eventhub_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Eventhub_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Eventhub_StatusARM instances for property testing - lazily instantiated by EventhubStatusARMGenerator()
var eventhubStatusARMGenerator gopter.Gen

// EventhubStatusARMGenerator returns a generator of Eventhub_StatusARM instances for property testing.
// We first initialize eventhubStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventhubStatusARMGenerator() gopter.Gen {
	if eventhubStatusARMGenerator != nil {
		return eventhubStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhubStatusARM(generators)
	eventhubStatusARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhubStatusARM(generators)
	AddRelatedPropertyGeneratorsForEventhubStatusARM(generators)
	eventhubStatusARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_StatusARM{}), generators)

	return eventhubStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForEventhubStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventhubStatusARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventhubStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventhubStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(EventhubStatusPropertiesARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemDataStatusARMGenerator())
}

func Test_Eventhub_Status_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Eventhub_Status_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForEventhubStatusPropertiesARM, EventhubStatusPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForEventhubStatusPropertiesARM runs a test to see if a specific instance of Eventhub_Status_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForEventhubStatusPropertiesARM(subject Eventhub_Status_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Eventhub_Status_PropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Eventhub_Status_PropertiesARM instances for property testing - lazily instantiated by
// EventhubStatusPropertiesARMGenerator()
var eventhubStatusPropertiesARMGenerator gopter.Gen

// EventhubStatusPropertiesARMGenerator returns a generator of Eventhub_Status_PropertiesARM instances for property testing.
// We first initialize eventhubStatusPropertiesARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func EventhubStatusPropertiesARMGenerator() gopter.Gen {
	if eventhubStatusPropertiesARMGenerator != nil {
		return eventhubStatusPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhubStatusPropertiesARM(generators)
	eventhubStatusPropertiesARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_Status_PropertiesARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForEventhubStatusPropertiesARM(generators)
	AddRelatedPropertyGeneratorsForEventhubStatusPropertiesARM(generators)
	eventhubStatusPropertiesARMGenerator = gen.Struct(reflect.TypeOf(Eventhub_Status_PropertiesARM{}), generators)

	return eventhubStatusPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForEventhubStatusPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForEventhubStatusPropertiesARM(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["MessageRetentionInDays"] = gen.PtrOf(gen.Int())
	gens["PartitionCount"] = gen.PtrOf(gen.Int())
	gens["PartitionIds"] = gen.SliceOf(gen.AlphaString())
	gens["Status"] = gen.PtrOf(gen.OneConstOf(
		EventhubStatusPropertiesStatusActive,
		EventhubStatusPropertiesStatusCreating,
		EventhubStatusPropertiesStatusDeleting,
		EventhubStatusPropertiesStatusDisabled,
		EventhubStatusPropertiesStatusReceiveDisabled,
		EventhubStatusPropertiesStatusRenaming,
		EventhubStatusPropertiesStatusRestoring,
		EventhubStatusPropertiesStatusSendDisabled,
		EventhubStatusPropertiesStatusUnknown))
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForEventhubStatusPropertiesARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForEventhubStatusPropertiesARM(gens map[string]gopter.Gen) {
	gens["CaptureDescription"] = gen.PtrOf(CaptureDescriptionStatusARMGenerator())
}

func Test_CaptureDescription_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of CaptureDescription_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForCaptureDescriptionStatusARM, CaptureDescriptionStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForCaptureDescriptionStatusARM runs a test to see if a specific instance of CaptureDescription_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForCaptureDescriptionStatusARM(subject CaptureDescription_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual CaptureDescription_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of CaptureDescription_StatusARM instances for property testing - lazily instantiated by
// CaptureDescriptionStatusARMGenerator()
var captureDescriptionStatusARMGenerator gopter.Gen

// CaptureDescriptionStatusARMGenerator returns a generator of CaptureDescription_StatusARM instances for property testing.
// We first initialize captureDescriptionStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func CaptureDescriptionStatusARMGenerator() gopter.Gen {
	if captureDescriptionStatusARMGenerator != nil {
		return captureDescriptionStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescriptionStatusARM(generators)
	captureDescriptionStatusARMGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForCaptureDescriptionStatusARM(generators)
	AddRelatedPropertyGeneratorsForCaptureDescriptionStatusARM(generators)
	captureDescriptionStatusARMGenerator = gen.Struct(reflect.TypeOf(CaptureDescription_StatusARM{}), generators)

	return captureDescriptionStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForCaptureDescriptionStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForCaptureDescriptionStatusARM(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
	gens["Encoding"] = gen.PtrOf(gen.OneConstOf(CaptureDescriptionStatusEncodingAvro, CaptureDescriptionStatusEncodingAvroDeflate))
	gens["IntervalInSeconds"] = gen.PtrOf(gen.Int())
	gens["SizeLimitInBytes"] = gen.PtrOf(gen.Int())
	gens["SkipEmptyArchives"] = gen.PtrOf(gen.Bool())
}

// AddRelatedPropertyGeneratorsForCaptureDescriptionStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForCaptureDescriptionStatusARM(gens map[string]gopter.Gen) {
	gens["Destination"] = gen.PtrOf(DestinationStatusARMGenerator())
}

func Test_Destination_StatusARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_StatusARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestinationStatusARM, DestinationStatusARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestinationStatusARM runs a test to see if a specific instance of Destination_StatusARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestinationStatusARM(subject Destination_StatusARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_StatusARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Destination_StatusARM instances for property testing - lazily instantiated by
// DestinationStatusARMGenerator()
var destinationStatusARMGenerator gopter.Gen

// DestinationStatusARMGenerator returns a generator of Destination_StatusARM instances for property testing.
// We first initialize destinationStatusARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DestinationStatusARMGenerator() gopter.Gen {
	if destinationStatusARMGenerator != nil {
		return destinationStatusARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestinationStatusARM(generators)
	destinationStatusARMGenerator = gen.Struct(reflect.TypeOf(Destination_StatusARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestinationStatusARM(generators)
	AddRelatedPropertyGeneratorsForDestinationStatusARM(generators)
	destinationStatusARMGenerator = gen.Struct(reflect.TypeOf(Destination_StatusARM{}), generators)

	return destinationStatusARMGenerator
}

// AddIndependentPropertyGeneratorsForDestinationStatusARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestinationStatusARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDestinationStatusARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDestinationStatusARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DestinationStatusPropertiesARMGenerator())
}

func Test_Destination_Status_PropertiesARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Destination_Status_PropertiesARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDestinationStatusPropertiesARM, DestinationStatusPropertiesARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDestinationStatusPropertiesARM runs a test to see if a specific instance of Destination_Status_PropertiesARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDestinationStatusPropertiesARM(subject Destination_Status_PropertiesARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Destination_Status_PropertiesARM
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Destination_Status_PropertiesARM instances for property testing - lazily instantiated by
// DestinationStatusPropertiesARMGenerator()
var destinationStatusPropertiesARMGenerator gopter.Gen

// DestinationStatusPropertiesARMGenerator returns a generator of Destination_Status_PropertiesARM instances for property testing.
func DestinationStatusPropertiesARMGenerator() gopter.Gen {
	if destinationStatusPropertiesARMGenerator != nil {
		return destinationStatusPropertiesARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDestinationStatusPropertiesARM(generators)
	destinationStatusPropertiesARMGenerator = gen.Struct(reflect.TypeOf(Destination_Status_PropertiesARM{}), generators)

	return destinationStatusPropertiesARMGenerator
}

// AddIndependentPropertyGeneratorsForDestinationStatusPropertiesARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDestinationStatusPropertiesARM(gens map[string]gopter.Gen) {
	gens["ArchiveNameFormat"] = gen.PtrOf(gen.AlphaString())
	gens["BlobContainer"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeAccountName"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeFolderPath"] = gen.PtrOf(gen.AlphaString())
	gens["DataLakeSubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["StorageAccountResourceId"] = gen.PtrOf(gen.AlphaString())
}