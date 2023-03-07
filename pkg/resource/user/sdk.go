// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package user

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/memorydb"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/memorydb-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.MemoryDB{}
	_ = &svcapitypes.User{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeUsersOutput
	resp, err = rm.sdkapi.DescribeUsersWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeUsers", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UserNotFoundFault" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.Users {
		if elem.ACLNames != nil {
			f0 := []*string{}
			for _, f0iter := range elem.ACLNames {
				var f0elem string
				f0elem = *f0iter
				f0 = append(f0, &f0elem)
			}
			ko.Status.ACLNames = f0
		} else {
			ko.Status.ACLNames = nil
		}
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.AccessString != nil {
			ko.Spec.AccessString = elem.AccessString
		} else {
			ko.Spec.AccessString = nil
		}
		if elem.Authentication != nil {
			f3 := &svcapitypes.Authentication{}
			if elem.Authentication.PasswordCount != nil {
				f3.PasswordCount = elem.Authentication.PasswordCount
			}
			if elem.Authentication.Type != nil {
				f3.Type = elem.Authentication.Type
			}
			ko.Status.Authentication = f3
		} else {
			ko.Status.Authentication = nil
		}
		if elem.MinimumEngineVersion != nil {
			ko.Status.MinimumEngineVersion = elem.MinimumEngineVersion
		} else {
			ko.Status.MinimumEngineVersion = nil
		}
		if elem.Name != nil {
			ko.Spec.Name = elem.Name
		} else {
			ko.Spec.Name = nil
		}
		if elem.Status != nil {
			ko.Status.Status = elem.Status
		} else {
			ko.Status.Status = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	if rm.isUserActive(&resource{ko}) {
		resourceARN := (*string)(ko.Status.ACKResourceMetadata.ARN)
		tags, err := rm.getTags(ctx, *resourceARN)
		if err != nil {
			return nil, err
		}
		ko.Spec.Tags = tags
	}

	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return r.ko.Spec.Name == nil

}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeUsersInput, error) {
	res := &svcsdk.DescribeUsersInput{}

	if r.ko.Spec.Name != nil {
		res.SetUserName(*r.ko.Spec.Name)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateUserOutput
	_ = resp
	resp, err = rm.sdkapi.CreateUserWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateUser", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.User.ACLNames != nil {
		f0 := []*string{}
		for _, f0iter := range resp.User.ACLNames {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		ko.Status.ACLNames = f0
	} else {
		ko.Status.ACLNames = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.User.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.User.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.User.AccessString != nil {
		ko.Spec.AccessString = resp.User.AccessString
	} else {
		ko.Spec.AccessString = nil
	}
	if resp.User.Authentication != nil {
		f3 := &svcapitypes.Authentication{}
		if resp.User.Authentication.PasswordCount != nil {
			f3.PasswordCount = resp.User.Authentication.PasswordCount
		}
		if resp.User.Authentication.Type != nil {
			f3.Type = resp.User.Authentication.Type
		}
		ko.Status.Authentication = f3
	} else {
		ko.Status.Authentication = nil
	}
	if resp.User.MinimumEngineVersion != nil {
		ko.Status.MinimumEngineVersion = resp.User.MinimumEngineVersion
	} else {
		ko.Status.MinimumEngineVersion = nil
	}
	if resp.User.Name != nil {
		ko.Spec.Name = resp.User.Name
	} else {
		ko.Spec.Name = nil
	}
	if resp.User.Status != nil {
		ko.Status.Status = resp.User.Status
	} else {
		ko.Status.Status = nil
	}

	rm.setStatusDefaults(ko)
	rm.setAnnotationsFields(desired, ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateUserInput, error) {
	res := &svcsdk.CreateUserInput{}

	if r.ko.Spec.AccessString != nil {
		res.SetAccessString(*r.ko.Spec.AccessString)
	}
	if r.ko.Spec.AuthenticationMode != nil {
		f1 := &svcsdk.AuthenticationMode{}
		if r.ko.Spec.AuthenticationMode.Passwords != nil {
			f1f0 := []*string{}
			for _, f1f0iter := range r.ko.Spec.AuthenticationMode.Passwords {
				var f1f0elem string
				if f1f0iter != nil {
					tmpSecret, err := rm.rr.SecretValueFromReference(ctx, f1f0iter)
					if err != nil {
						return nil, ackrequeue.Needed(err)
					}
					if tmpSecret != "" {
						f1f0elem = tmpSecret
					}
				}
				f1f0 = append(f1f0, &f1f0elem)
			}
			f1.SetPasswords(f1f0)
		}
		if r.ko.Spec.AuthenticationMode.Type != nil {
			f1.SetType(*r.ko.Spec.AuthenticationMode.Type)
		}
		res.SetAuthenticationMode(f1)
	}
	if r.ko.Spec.Tags != nil {
		f2 := []*svcsdk.Tag{}
		for _, f2iter := range r.ko.Spec.Tags {
			f2elem := &svcsdk.Tag{}
			if f2iter.Key != nil {
				f2elem.SetKey(*f2iter.Key)
			}
			if f2iter.Value != nil {
				f2elem.SetValue(*f2iter.Value)
			}
			f2 = append(f2, f2elem)
		}
		res.SetTags(f2)
	}
	if r.ko.Spec.Name != nil {
		res.SetUserName(*r.ko.Spec.Name)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	res, err := rm.validateUserNeedsUpdate(desired, latest, delta)

	if err != nil || res != nil {
		return res, err
	}

	if delta.DifferentAt("Spec.Tags") {
		err = rm.updateTags(ctx, desired, latest)
		if err != nil {
			return nil, err
		}
	}

	if !delta.DifferentExcept("Spec.Tags") {
		return desired, nil
	}

	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdateUserOutput
	_ = resp
	resp, err = rm.sdkapi.UpdateUserWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdateUser", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.User.ACLNames != nil {
		f0 := []*string{}
		for _, f0iter := range resp.User.ACLNames {
			var f0elem string
			f0elem = *f0iter
			f0 = append(f0, &f0elem)
		}
		ko.Status.ACLNames = f0
	} else {
		ko.Status.ACLNames = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.User.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.User.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.User.AccessString != nil {
		ko.Spec.AccessString = resp.User.AccessString
	} else {
		ko.Spec.AccessString = nil
	}
	if resp.User.Authentication != nil {
		f3 := &svcapitypes.Authentication{}
		if resp.User.Authentication.PasswordCount != nil {
			f3.PasswordCount = resp.User.Authentication.PasswordCount
		}
		if resp.User.Authentication.Type != nil {
			f3.Type = resp.User.Authentication.Type
		}
		ko.Status.Authentication = f3
	} else {
		ko.Status.Authentication = nil
	}
	if resp.User.MinimumEngineVersion != nil {
		ko.Status.MinimumEngineVersion = resp.User.MinimumEngineVersion
	} else {
		ko.Status.MinimumEngineVersion = nil
	}
	if resp.User.Name != nil {
		ko.Spec.Name = resp.User.Name
	} else {
		ko.Spec.Name = nil
	}
	if resp.User.Status != nil {
		ko.Status.Status = resp.User.Status
	} else {
		ko.Status.Status = nil
	}

	rm.setStatusDefaults(ko)
	rm.setAnnotationsFields(desired, ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdateUserInput, error) {
	res := &svcsdk.UpdateUserInput{}

	if r.ko.Spec.AccessString != nil {
		res.SetAccessString(*r.ko.Spec.AccessString)
	}
	if r.ko.Spec.AuthenticationMode != nil {
		f1 := &svcsdk.AuthenticationMode{}
		if r.ko.Spec.AuthenticationMode.Passwords != nil {
			f1f0 := []*string{}
			for _, f1f0iter := range r.ko.Spec.AuthenticationMode.Passwords {
				var f1f0elem string
				if f1f0iter != nil {
					tmpSecret, err := rm.rr.SecretValueFromReference(ctx, f1f0iter)
					if err != nil {
						return nil, ackrequeue.Needed(err)
					}
					if tmpSecret != "" {
						f1f0elem = tmpSecret
					}
				}
				f1f0 = append(f1f0, &f1f0elem)
			}
			f1.SetPasswords(f1f0)
		}
		if r.ko.Spec.AuthenticationMode.Type != nil {
			f1.SetType(*r.ko.Spec.AuthenticationMode.Type)
		}
		res.SetAuthenticationMode(f1)
	}
	if r.ko.Spec.Name != nil {
		res.SetUserName(*r.ko.Spec.Name)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteUserOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteUserWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteUser", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteUserInput, error) {
	res := &svcsdk.DeleteUserInput{}

	if r.ko.Spec.Name != nil {
		res.SetUserName(*r.ko.Spec.Name)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.User,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "UserAlreadyExistsFault",
		"InvalidParameterValueException",
		"DuplicateUserNameFault",
		"InvalidParameterCombinationException":
		return true
	default:
		return false
	}
}
