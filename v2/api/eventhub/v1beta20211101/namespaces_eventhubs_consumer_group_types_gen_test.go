// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"encoding/json"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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

func Test_NamespacesEventhubsConsumerGroup_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup to hub returns original",
		prop.ForAll(RunResourceConversionTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNamespacesEventhubsConsumerGroup tests if a specific instance of NamespacesEventhubsConsumerGroup round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20211101s.NamespacesEventhubsConsumerGroup
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NamespacesEventhubsConsumerGroup
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubsConsumerGroup_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumerGroup to NamespacesEventhubsConsumerGroup via AssignPropertiesToNamespacesEventhubsConsumerGroup & AssignPropertiesFromNamespacesEventhubsConsumerGroup returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup tests if a specific instance of NamespacesEventhubsConsumerGroup can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.NamespacesEventhubsConsumerGroup
	err := copied.AssignPropertiesToNamespacesEventhubsConsumerGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsConsumerGroup
	err = actual.AssignPropertiesFromNamespacesEventhubsConsumerGroup(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubsConsumerGroup_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumerGroup via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumerGroup, NamespacesEventhubsConsumerGroupGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumerGroup runs a test to see if a specific instance of NamespacesEventhubsConsumerGroup round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumerGroup(subject NamespacesEventhubsConsumerGroup) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumerGroup
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

// Generator of NamespacesEventhubsConsumerGroup instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumerGroupGenerator()
var namespacesEventhubsConsumerGroupGenerator gopter.Gen

// NamespacesEventhubsConsumerGroupGenerator returns a generator of NamespacesEventhubsConsumerGroup instances for property testing.
func NamespacesEventhubsConsumerGroupGenerator() gopter.Gen {
	if namespacesEventhubsConsumerGroupGenerator != nil {
		return namespacesEventhubsConsumerGroupGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup(generators)
	namespacesEventhubsConsumerGroupGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumerGroup{}), generators)

	return namespacesEventhubsConsumerGroupGenerator
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsConsumerGroup(gens map[string]gopter.Gen) {
	gens["Spec"] = NamespacesEventhubsConsumergroupsSpecGenerator()
	gens["Status"] = ConsumerGroupSTATUSGenerator()
}

func Test_ConsumerGroup_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ConsumerGroup_STATUS to ConsumerGroup_STATUS via AssignPropertiesToConsumerGroupSTATUS & AssignPropertiesFromConsumerGroupSTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForConsumerGroupSTATUS, ConsumerGroupSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForConsumerGroupSTATUS tests if a specific instance of ConsumerGroup_STATUS can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForConsumerGroupSTATUS(subject ConsumerGroup_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.ConsumerGroup_STATUS
	err := copied.AssignPropertiesToConsumerGroupSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ConsumerGroup_STATUS
	err = actual.AssignPropertiesFromConsumerGroupSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ConsumerGroup_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ConsumerGroup_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForConsumerGroupSTATUS, ConsumerGroupSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForConsumerGroupSTATUS runs a test to see if a specific instance of ConsumerGroup_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForConsumerGroupSTATUS(subject ConsumerGroup_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ConsumerGroup_STATUS
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

// Generator of ConsumerGroup_STATUS instances for property testing - lazily instantiated by
// ConsumerGroupSTATUSGenerator()
var consumerGroupSTATUSGenerator gopter.Gen

// ConsumerGroupSTATUSGenerator returns a generator of ConsumerGroup_STATUS instances for property testing.
// We first initialize consumerGroupSTATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ConsumerGroupSTATUSGenerator() gopter.Gen {
	if consumerGroupSTATUSGenerator != nil {
		return consumerGroupSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroupSTATUS(generators)
	consumerGroupSTATUSGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForConsumerGroupSTATUS(generators)
	AddRelatedPropertyGeneratorsForConsumerGroupSTATUS(generators)
	consumerGroupSTATUSGenerator = gen.Struct(reflect.TypeOf(ConsumerGroup_STATUS{}), generators)

	return consumerGroupSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForConsumerGroupSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForConsumerGroupSTATUS(gens map[string]gopter.Gen) {
	gens["CreatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedAt"] = gen.PtrOf(gen.AlphaString())
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForConsumerGroupSTATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForConsumerGroupSTATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemDataSTATUSGenerator())
}

func Test_NamespacesEventhubsConsumergroups_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NamespacesEventhubsConsumergroups_Spec to NamespacesEventhubsConsumergroups_Spec via AssignPropertiesToNamespacesEventhubsConsumergroupsSpec & AssignPropertiesFromNamespacesEventhubsConsumergroupsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNamespacesEventhubsConsumergroupsSpec, NamespacesEventhubsConsumergroupsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNamespacesEventhubsConsumergroupsSpec tests if a specific instance of NamespacesEventhubsConsumergroups_Spec can be assigned to v1beta20211101storage and back losslessly
func RunPropertyAssignmentTestForNamespacesEventhubsConsumergroupsSpec(subject NamespacesEventhubsConsumergroups_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20211101s.NamespacesEventhubsConsumergroups_Spec
	err := copied.AssignPropertiesToNamespacesEventhubsConsumergroupsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NamespacesEventhubsConsumergroups_Spec
	err = actual.AssignPropertiesFromNamespacesEventhubsConsumergroupsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NamespacesEventhubsConsumergroups_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsConsumergroups_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsConsumergroupsSpec, NamespacesEventhubsConsumergroupsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsConsumergroupsSpec runs a test to see if a specific instance of NamespacesEventhubsConsumergroups_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsConsumergroupsSpec(subject NamespacesEventhubsConsumergroups_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsConsumergroups_Spec
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

// Generator of NamespacesEventhubsConsumergroups_Spec instances for property testing - lazily instantiated by
// NamespacesEventhubsConsumergroupsSpecGenerator()
var namespacesEventhubsConsumergroupsSpecGenerator gopter.Gen

// NamespacesEventhubsConsumergroupsSpecGenerator returns a generator of NamespacesEventhubsConsumergroups_Spec instances for property testing.
func NamespacesEventhubsConsumergroupsSpecGenerator() gopter.Gen {
	if namespacesEventhubsConsumergroupsSpecGenerator != nil {
		return namespacesEventhubsConsumergroupsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroupsSpec(generators)
	namespacesEventhubsConsumergroupsSpecGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsConsumergroups_Spec{}), generators)

	return namespacesEventhubsConsumergroupsSpecGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroupsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsConsumergroupsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["UserMetadata"] = gen.PtrOf(gen.AlphaString())
}