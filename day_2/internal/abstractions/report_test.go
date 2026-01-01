package abstractions

import "testing"

func TestReport_GetStatus(t *testing.T) {
	tests := []struct {
		name           string
		levels         []Level
		expectedStatus Status
	}{
		{
			name: "decreasing_by_1_or_2_is_safe (7 6 4 2 1)",
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
			name: "increase_jump_greater_than_3_is_unsafe (1 2 7 8 9)",
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
			name: "decrease_drop_greater_than_3_is_unsafe (9 7 6 2 1)",
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
			name: "direction_change_is_unsafe (1 3 2 4 5)",
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
			name: "equal_adjacent_levels_is_unsafe (8 6 4 4 1)",
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
			name: "increasing_by_1_to_3_is_safe (1 3 6 7 9)",
			levels: []Level{
				Level(1),
				Level(3),
				Level(6),
				Level(7),
				Level(9),
			},
			expectedStatus: StatusSafe,
		},
		{
			name: "direction_change_late_is_unsafe (89 91 92 95 93 94)",
			levels: []Level{
				Level(89),
				Level(91),
				Level(92),
				Level(95),
				Level(93),
				Level(94),
			},
			expectedStatus: StatusSafe,
		},
		{
			name: "direction_change_with_dip_is_unsafe (77 78 79 82 79 83)",
			levels: []Level{
				Level(77),
				Level(78),
				Level(79),
				Level(82),
				Level(79),
				Level(83),
			},
			expectedStatus: StatusSafe,
		},
		{
			name: "direction_change_early_is_unsafe (55 54 57 60 61)",
			levels: []Level{
				Level(55),
				Level(54),
				Level(57),
				Level(60),
				Level(61),
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
