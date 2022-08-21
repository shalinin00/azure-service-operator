// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201storage

import (
	"encoding/json"
	v20210301s "github.com/Azure/azure-service-operator/v2/api/cache/v1beta20210301storage"
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

func Test_Redis_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedis, RedisGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedis runs a test to see if a specific instance of Redis round trips to JSON and back losslessly
func RunJSONSerializationTestForRedis(subject Redis) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis
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

// Generator of Redis instances for property testing - lazily instantiated by RedisGenerator()
var redisGenerator gopter.Gen

// RedisGenerator returns a generator of Redis instances for property testing.
func RedisGenerator() gopter.Gen {
	if redisGenerator != nil {
		return redisGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedis(generators)
	redisGenerator = gen.Struct(reflect.TypeOf(Redis{}), generators)

	return redisGenerator
}

// AddRelatedPropertyGeneratorsForRedis is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedis(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisSpecGenerator()
	gens["Status"] = RedisResourceSTATUSGenerator()
}

func Test_Redis_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Redis_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisSpec, RedisSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisSpec runs a test to see if a specific instance of Redis_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisSpec(subject Redis_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Redis_Spec
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

// Generator of Redis_Spec instances for property testing - lazily instantiated by RedisSpecGenerator()
var redisSpecGenerator gopter.Gen

// RedisSpecGenerator returns a generator of Redis_Spec instances for property testing.
// We first initialize redisSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisSpecGenerator() gopter.Gen {
	if redisSpecGenerator != nil {
		return redisSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisSpec(generators)
	redisSpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisSpec(generators)
	AddRelatedPropertyGeneratorsForRedisSpec(generators)
	redisSpecGenerator = gen.Struct(reflect.TypeOf(Redis_Spec{}), generators)

	return redisSpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["RedisConfiguration"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisSpec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(RedisOperatorSpecGenerator())
	gens["Sku"] = gen.PtrOf(SkuGenerator())
}

func Test_RedisResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisResourceSTATUS, RedisResourceSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisResourceSTATUS runs a test to see if a specific instance of RedisResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisResourceSTATUS(subject RedisResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisResource_STATUS
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

// Generator of RedisResource_STATUS instances for property testing - lazily instantiated by
// RedisResourceSTATUSGenerator()
var redisResourceSTATUSGenerator gopter.Gen

// RedisResourceSTATUSGenerator returns a generator of RedisResource_STATUS instances for property testing.
// We first initialize redisResourceSTATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisResourceSTATUSGenerator() gopter.Gen {
	if redisResourceSTATUSGenerator != nil {
		return redisResourceSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceSTATUS(generators)
	redisResourceSTATUSGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisResourceSTATUS(generators)
	AddRelatedPropertyGeneratorsForRedisResourceSTATUS(generators)
	redisResourceSTATUSGenerator = gen.Struct(reflect.TypeOf(RedisResource_STATUS{}), generators)

	return redisResourceSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisResourceSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisResourceSTATUS(gens map[string]gopter.Gen) {
	gens["EnableNonSslPort"] = gen.PtrOf(gen.Bool())
	gens["HostName"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["MinimumTlsVersion"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["PublicNetworkAccess"] = gen.PtrOf(gen.AlphaString())
	gens["RedisConfiguration"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["RedisVersion"] = gen.PtrOf(gen.AlphaString())
	gens["ReplicasPerMaster"] = gen.PtrOf(gen.Int())
	gens["ReplicasPerPrimary"] = gen.PtrOf(gen.Int())
	gens["ShardCount"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["StaticIP"] = gen.PtrOf(gen.AlphaString())
	gens["SubnetId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["TenantSettings"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisResourceSTATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisResourceSTATUS(gens map[string]gopter.Gen) {
	gens["Instances"] = gen.SliceOf(RedisInstanceDetailsSTATUSGenerator())
	gens["LinkedServers"] = gen.SliceOf(RedisLinkedServerSTATUSGenerator())
	gens["PrivateEndpointConnections"] = gen.SliceOf(PrivateEndpointConnectionSTATUSSubResourceEmbeddedGenerator())
	gens["Sku"] = gen.PtrOf(SkuSTATUSGenerator())
}

func Test_PrivateEndpointConnection_STATUS_SubResourceEmbedded_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from PrivateEndpointConnection_STATUS_SubResourceEmbedded to PrivateEndpointConnection_STATUS_SubResourceEmbedded via AssignPropertiesToPrivateEndpointConnectionSTATUSSubResourceEmbedded & AssignPropertiesFromPrivateEndpointConnectionSTATUSSubResourceEmbedded returns original",
		prop.ForAll(RunPropertyAssignmentTestForPrivateEndpointConnectionSTATUSSubResourceEmbedded, PrivateEndpointConnectionSTATUSSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForPrivateEndpointConnectionSTATUSSubResourceEmbedded tests if a specific instance of PrivateEndpointConnection_STATUS_SubResourceEmbedded can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForPrivateEndpointConnectionSTATUSSubResourceEmbedded(subject PrivateEndpointConnection_STATUS_SubResourceEmbedded) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.PrivateEndpointConnection_STATUS_SubResourceEmbedded
	err := copied.AssignPropertiesToPrivateEndpointConnectionSTATUSSubResourceEmbedded(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual PrivateEndpointConnection_STATUS_SubResourceEmbedded
	err = actual.AssignPropertiesFromPrivateEndpointConnectionSTATUSSubResourceEmbedded(&other)
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

func Test_PrivateEndpointConnection_STATUS_SubResourceEmbedded_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointConnection_STATUS_SubResourceEmbedded via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointConnectionSTATUSSubResourceEmbedded, PrivateEndpointConnectionSTATUSSubResourceEmbeddedGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointConnectionSTATUSSubResourceEmbedded runs a test to see if a specific instance of PrivateEndpointConnection_STATUS_SubResourceEmbedded round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointConnectionSTATUSSubResourceEmbedded(subject PrivateEndpointConnection_STATUS_SubResourceEmbedded) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointConnection_STATUS_SubResourceEmbedded
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

// Generator of PrivateEndpointConnection_STATUS_SubResourceEmbedded instances for property testing - lazily
// instantiated by PrivateEndpointConnectionSTATUSSubResourceEmbeddedGenerator()
var privateEndpointConnectionSTATUSSubResourceEmbeddedGenerator gopter.Gen

// PrivateEndpointConnectionSTATUSSubResourceEmbeddedGenerator returns a generator of PrivateEndpointConnection_STATUS_SubResourceEmbedded instances for property testing.
func PrivateEndpointConnectionSTATUSSubResourceEmbeddedGenerator() gopter.Gen {
	if privateEndpointConnectionSTATUSSubResourceEmbeddedGenerator != nil {
		return privateEndpointConnectionSTATUSSubResourceEmbeddedGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSSubResourceEmbedded(generators)
	privateEndpointConnectionSTATUSSubResourceEmbeddedGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointConnection_STATUS_SubResourceEmbedded{}), generators)

	return privateEndpointConnectionSTATUSSubResourceEmbeddedGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSSubResourceEmbedded is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointConnectionSTATUSSubResourceEmbedded(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisInstanceDetails_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisInstanceDetails_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisInstanceDetailsSTATUS, RedisInstanceDetailsSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisInstanceDetailsSTATUS runs a test to see if a specific instance of RedisInstanceDetails_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisInstanceDetailsSTATUS(subject RedisInstanceDetails_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisInstanceDetails_STATUS
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

// Generator of RedisInstanceDetails_STATUS instances for property testing - lazily instantiated by
// RedisInstanceDetailsSTATUSGenerator()
var redisInstanceDetailsSTATUSGenerator gopter.Gen

// RedisInstanceDetailsSTATUSGenerator returns a generator of RedisInstanceDetails_STATUS instances for property testing.
func RedisInstanceDetailsSTATUSGenerator() gopter.Gen {
	if redisInstanceDetailsSTATUSGenerator != nil {
		return redisInstanceDetailsSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisInstanceDetailsSTATUS(generators)
	redisInstanceDetailsSTATUSGenerator = gen.Struct(reflect.TypeOf(RedisInstanceDetails_STATUS{}), generators)

	return redisInstanceDetailsSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisInstanceDetailsSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisInstanceDetailsSTATUS(gens map[string]gopter.Gen) {
	gens["IsMaster"] = gen.PtrOf(gen.Bool())
	gens["IsPrimary"] = gen.PtrOf(gen.Bool())
	gens["NonSslPort"] = gen.PtrOf(gen.Int())
	gens["ShardId"] = gen.PtrOf(gen.Int())
	gens["SslPort"] = gen.PtrOf(gen.Int())
	gens["Zone"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisLinkedServer_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisLinkedServer_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisLinkedServerSTATUS, RedisLinkedServerSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisLinkedServerSTATUS runs a test to see if a specific instance of RedisLinkedServer_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisLinkedServerSTATUS(subject RedisLinkedServer_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisLinkedServer_STATUS
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

// Generator of RedisLinkedServer_STATUS instances for property testing - lazily instantiated by
// RedisLinkedServerSTATUSGenerator()
var redisLinkedServerSTATUSGenerator gopter.Gen

// RedisLinkedServerSTATUSGenerator returns a generator of RedisLinkedServer_STATUS instances for property testing.
func RedisLinkedServerSTATUSGenerator() gopter.Gen {
	if redisLinkedServerSTATUSGenerator != nil {
		return redisLinkedServerSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisLinkedServerSTATUS(generators)
	redisLinkedServerSTATUSGenerator = gen.Struct(reflect.TypeOf(RedisLinkedServer_STATUS{}), generators)

	return redisLinkedServerSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisLinkedServerSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisLinkedServerSTATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisOperatorSpec, RedisOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisOperatorSpec runs a test to see if a specific instance of RedisOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisOperatorSpec(subject RedisOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisOperatorSpec
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

// Generator of RedisOperatorSpec instances for property testing - lazily instantiated by RedisOperatorSpecGenerator()
var redisOperatorSpecGenerator gopter.Gen

// RedisOperatorSpecGenerator returns a generator of RedisOperatorSpec instances for property testing.
func RedisOperatorSpecGenerator() gopter.Gen {
	if redisOperatorSpecGenerator != nil {
		return redisOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisOperatorSpec(generators)
	redisOperatorSpecGenerator = gen.Struct(reflect.TypeOf(RedisOperatorSpec{}), generators)

	return redisOperatorSpecGenerator
}

// AddRelatedPropertyGeneratorsForRedisOperatorSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisOperatorSpec(gens map[string]gopter.Gen) {
	gens["Secrets"] = gen.PtrOf(RedisOperatorSecretsGenerator())
}

func Test_Sku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku to Sku via AssignPropertiesToSku & AssignPropertiesFromSku returns original",
		prop.ForAll(RunPropertyAssignmentTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSku tests if a specific instance of Sku can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForSku(subject Sku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Sku
	err := copied.AssignPropertiesToSku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku
	err = actual.AssignPropertiesFromSku(&other)
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

func Test_Sku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSku, SkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSku runs a test to see if a specific instance of Sku round trips to JSON and back losslessly
func RunJSONSerializationTestForSku(subject Sku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku
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

// Generator of Sku instances for property testing - lazily instantiated by SkuGenerator()
var skuGenerator gopter.Gen

// SkuGenerator returns a generator of Sku instances for property testing.
func SkuGenerator() gopter.Gen {
	if skuGenerator != nil {
		return skuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSku(generators)
	skuGenerator = gen.Struct(reflect.TypeOf(Sku{}), generators)

	return skuGenerator
}

// AddIndependentPropertyGeneratorsForSku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSku(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_Sku_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Sku_STATUS to Sku_STATUS via AssignPropertiesToSkuSTATUS & AssignPropertiesFromSkuSTATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSkuSTATUS, SkuSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSkuSTATUS tests if a specific instance of Sku_STATUS can be assigned to v1beta20210301storage and back losslessly
func RunPropertyAssignmentTestForSkuSTATUS(subject Sku_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20210301s.Sku_STATUS
	err := copied.AssignPropertiesToSkuSTATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Sku_STATUS
	err = actual.AssignPropertiesFromSkuSTATUS(&other)
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

func Test_Sku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Sku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSkuSTATUS, SkuSTATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSkuSTATUS runs a test to see if a specific instance of Sku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSkuSTATUS(subject Sku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Sku_STATUS
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

// Generator of Sku_STATUS instances for property testing - lazily instantiated by SkuSTATUSGenerator()
var skuSTATUSGenerator gopter.Gen

// SkuSTATUSGenerator returns a generator of Sku_STATUS instances for property testing.
func SkuSTATUSGenerator() gopter.Gen {
	if skuSTATUSGenerator != nil {
		return skuSTATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSkuSTATUS(generators)
	skuSTATUSGenerator = gen.Struct(reflect.TypeOf(Sku_STATUS{}), generators)

	return skuSTATUSGenerator
}

// AddIndependentPropertyGeneratorsForSkuSTATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSkuSTATUS(gens map[string]gopter.Gen) {
	gens["Capacity"] = gen.PtrOf(gen.Int())
	gens["Family"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

func Test_RedisOperatorSecrets_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisOperatorSecrets via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisOperatorSecrets, RedisOperatorSecretsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisOperatorSecrets runs a test to see if a specific instance of RedisOperatorSecrets round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisOperatorSecrets(subject RedisOperatorSecrets) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisOperatorSecrets
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

// Generator of RedisOperatorSecrets instances for property testing - lazily instantiated by
// RedisOperatorSecretsGenerator()
var redisOperatorSecretsGenerator gopter.Gen

// RedisOperatorSecretsGenerator returns a generator of RedisOperatorSecrets instances for property testing.
func RedisOperatorSecretsGenerator() gopter.Gen {
	if redisOperatorSecretsGenerator != nil {
		return redisOperatorSecretsGenerator
	}

	generators := make(map[string]gopter.Gen)
	redisOperatorSecretsGenerator = gen.Struct(reflect.TypeOf(RedisOperatorSecrets{}), generators)

	return redisOperatorSecretsGenerator
}