package abstractions

import "testing"

func TestReport_GetStatus(t *testing.T) {
	tests := []struct {
		name           string
		levels         []Level
		expectedStatus Status
	}{
		{
			name: "given ",
			levels: []Level{
				Level(7),
				Level(6),
				Level(4),
				Level(2),
				Level(1),
			},
			expectedStatus: StatusSafe,
		},
		{
			name: "given ",
			levels: []Level{
				Level(1),
				Level(2),
				Level(7),
				Level(8),
				Level(9),
			},
			expectedStatus: StatusUnsafe,
		},
		{
			name: "given ",
			levels: []Level{
				Level(9),
				Level(7),
				Level(6),
				Level(2),
				Level(1),
			},
			expectedStatus: StatusUnsafe,
		},
		{
			name: "given ",
			levels: []Level{
				Level(1),
				Level(3),
				Level(2),
				Level(4),
				Level(5),
			},
			expectedStatus: StatusSafe,
		},
		{
			name: "given ",
			levels: []Level{
				Level(8),
				Level(6),
				Level(4),
				Level(4),
				Level(1),
			},
			expectedStatus: StatusSafe,
		},
		{
			name: "given ",
			levels: []Level{
				Level(1),
				Level(3),
				Level(6),
				Level(7),
				Level(9),
			},
			expectedStatus: StatusSafe,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewReport(1, tt.levels).GetStatus()

			if got != tt.expectedStatus {
				t.Errorf("Report.GetStatus() = %v, want %v", got, tt.expectedStatus)
			}
		})
	}
}
