package gordle

import "testing"

func Test_feedback_String(t *testing.T) {
	tt := map[string]struct {
		fb   feedback
		want string
	}{
		"three correct": {
			fb:   feedback{correctPosition, correctPosition, correctPosition},
			want: "💚💚💚",
		},
		"one of each": {
			fb:   feedback{correctPosition, wrongPosition, absentCharacter},
			want: "💚🟡⬜️",
		},
		"different order for one of each": {
			fb:   feedback{wrongPosition, absentCharacter, correctPosition},
			want: "🟡⬜️💚",
		},
		"unknown position": {
			fb:   feedback{100},
			want: "💔",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			if got := tc.fb.String(); got != tc.want {
				t.Errorf("String() = %v, want %v", got, tc.want)
			}
		})
	}
}
