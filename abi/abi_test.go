// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package abi

import (
	"math/rand"
	"strings"
	"testing"

	spb "github.com/google/go-sev-guest/proto/sevsnp"
	"google.golang.org/protobuf/encoding/prototext"
)

var emptyReport = `
	version: 2
	policy: 0xa0000
	signature_algo: 1
	report_data: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01'
	family_id: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
  image_id: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	measurement: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
  host_data: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
  id_key_digest: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
  author_key_digest: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
  report_id: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
  report_id_ma: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
  chip_id: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
 	signature: '\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00'
	`

func TestReportMbz(t *testing.T) {
	tests := []struct {
		name        string
		changeIndex int
		changeValue byte
		wantErr     string
	}{
		{
			name:        "AuthorKeyEn reserved",
			changeIndex: 0x49,
			wantErr:     "mbz bits at offset 0x48 not zero: 0x0000cc00",
		},
		{
			name:        "pre-report data",
			changeIndex: 0x4f,
			wantErr:     "mbz range [0x4c:0x50] not all zero: 000000cc",
		},
		{
			name:        "pre-chip id",
			changeIndex: 0x18A,
			wantErr:     "mbz range [0x188:0x1a0] not all zero: 0000cc",
		},
		{
			name:        "current reserved",
			changeIndex: 0x1EB,
			wantErr:     "mbz range [0x1eb:0x1ec] not all zero: cc",
		},
		{
			name:        "committed reserved",
			changeIndex: 0x1EF,
			wantErr:     "mbz range [0x1ef:0x1f0] not all zero: cc",
		},
		{
			name:        "pre-signature reserved",
			changeIndex: 0x1f9,
			wantErr:     "mbz range [0x1f8:0x2a0] not all zero: 00cc",
		},
		{
			name:        "post-ecdsa signature reserved",
			changeIndex: signatureOffset + EcdsaP384Sha384SignatureSize + 2,
			wantErr:     "mbz range [0x330:0x4a0] not all zero: 0000cc",
		},
		{
			name:        "Guest policy bit 17",
			changeIndex: policyOffset + 2, // Bits 16-23
			changeValue: 0x1d,             // Set bits 16, 18, 19, 20
			wantErr:     "policy[17] is reserved, must be 1, got 0",
		},
		{
			name:        "Guest policy bit 21",
			changeIndex: policyOffset + 2, // Bits 16-23
			changeValue: 0x22,             // Set bits 17, 21
			wantErr:     "policy[63:21] are reserved mbz, got 0x220000",
		},
	}
	reportProto := &spb.Report{}
	if err := prototext.Unmarshal([]byte(emptyReport), reportProto); err != nil {
		t.Fatalf("test failure: %v", err)
	}
	for _, tc := range tests {
		// Everything but the signature hase
		raw, err := ReportToAbiBytes(reportProto)
		if err != nil {
			t.Fatalf("%s: test failure: ReportToAbiBytes(%v) errored unexpectedly: %v", tc.name, reportProto, err)
		}
		changeValue := byte(0xcc)
		if tc.changeValue != 0 {
			changeValue = tc.changeValue
		}
		raw[tc.changeIndex] = changeValue
		if _, err := ReportToProto(raw); !strings.Contains(err.Error(), tc.wantErr) {
			t.Errorf("%s: ReportToProto(%v) = _, %v. Want error %v", tc.name, reportProto, err, tc.wantErr)
		}
	}
}

func TestSnpPolicySection(t *testing.T) {
	entropySize := 128
	entropy := make([]uint8, entropySize)
	rand.Read(entropy)
	for tc := 0; tc < entropySize/3; tc++ {
		policy := SnpPolicy{
			ABIMinor:     entropy[tc*3],
			ABIMajor:     entropy[tc*3+1],
			SMT:          (entropy[tc*3+2] & 1) != 0,
			MigrateMA:    (entropy[tc*3+2] & 2) != 0,
			Debug:        (entropy[tc*3+2] & 4) != 0,
			SingleSocket: (entropy[tc*3+2] & 8) != 0,
		}

		got, err := ParseSnpPolicy(SnpPolicyToBytes(policy))
		if err != nil {
			t.Errorf("ParseSnpPolicy(SnpPolicyToBytes(%v)) errored unexpectedly: %v", policy, err)
		}
		if got != policy {
			t.Errorf("ParseSnpPolicy(SnpPolicyToBytes(%v)) = %v, want %v", policy, got, policy)
		}
	}
}

func TestSnpPlatformInfo(t *testing.T) {
	tests := []struct {
		input   uint64
		want    SnpPlatformInfo
		wantErr string
	}{
		{
			input: 0,
		},
		{
			input: 3,
			want:  SnpPlatformInfo{TSMEEnabled: true, SMTEnabled: true},
		},
		{
			input:   4,
			wantErr: "unrecognized platform info bit(s): 0x4",
		},
	}
	for _, tc := range tests {
		got, err := ParseSnpPlatformInfo(tc.input)
		if (err != nil && (tc.wantErr == "" || !strings.Contains(err.Error(), tc.wantErr))) ||
			(err == nil && tc.wantErr != "") {
			t.Errorf("ParseSnpPlatformInfo(%x) errored unexpectedly. Got %v, want %v",
				tc.input, err, tc.wantErr)
		}
		if err == nil && tc.want != got {
			t.Errorf("ParseSnpPlatformInfo(%x) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestDecomposeAuthorKeyEn(t *testing.T) {
	tests := []struct {
		input   AuthorKeyEnContainer
		want    AuthorKeyEnParts
		wantErr bool
	}{
		{
			input: 0,
			want:  AuthorKeyEnParts{SigningKey: SigningKeyVCEK},
		},
		{
			input: 0x1,
			want:  AuthorKeyEnParts{AuthorKeyEn: true, SigningKey: SigningKeyVCEK},
		},
		{
			input: 0x2,
			want:  AuthorKeyEnParts{MaskChipKey: true, SigningKey: SigningKeyVCEK},
		},
		{
			input: 0x3,
			want: AuthorKeyEnParts{
				AuthorKeyEn: true,
				MaskChipKey: true,
				SigningKey:  SigningKeyVCEK,
			},
		},
		{
			input: 0x4,
			want:  AuthorKeyEnParts{SigningKey: SigningKeyVLEK},
		},
		{
			input: 0x5,
			want:  AuthorKeyEnParts{AuthorKeyEn: true, SigningKey: SigningKeyVLEK},
		},
		{
			input: 0x6,
			want:  AuthorKeyEnParts{MaskChipKey: true, SigningKey: SigningKeyVLEK},
		},
		{
			input: 0x7,
			want: AuthorKeyEnParts{
				AuthorKeyEn: true,
				MaskChipKey: true,
				SigningKey:  SigningKeyVLEK,
			},
		},
		{
			input:   0x8,
			wantErr: true,
		},
		{
			input:   0x10,
			wantErr: true,
		},
		{
			input: 0x1c,
			want:  AuthorKeyEnParts{SigningKey: SigningKeyNone},
		},
		{
			input: 0x1f,
			want: AuthorKeyEnParts{
				AuthorKeyEn: true,
				MaskChipKey: true,
				SigningKey:  SigningKeyNone,
			},
		},
		{
			input:   0x20,
			wantErr: true,
		},
		{
			input:   0x80000000,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		got, err := DecomposeAuthorKeyEn(tc.input)

		if tc.wantErr != (err != nil) {
			if err == nil {
				t.Errorf("DecomposeAuthorKeyEn(%#x) did not error as expected", tc.input)
				continue
			}
			t.Errorf("DecomposeAuthorKeyEn(%#x) errored unexpectedly: %v", tc.input, err)
			continue
		}

		if got != tc.want {
			t.Errorf("DecomposeAuthorKeyEn(%#x) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
