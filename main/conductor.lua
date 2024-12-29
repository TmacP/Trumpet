local conductor = {}

-- Constants
conductor.crotchets_per_bar = 8
conductor.bpm = 120
conductor.offset = 0.2  -- Positive means the song must be minused!
conductor.offset_static = 0.4

-- State Variables
conductor.song_position = 0
conductor.delta_song_pos = 0
conductor.last_hit = 0  -- Last (snapped to beat) time spacebar was pressed
conductor.actual_last_hit = 0
conductor.next_beat_time = 0
conductor.next_bar_time = 0
conductor.add_offset = 0
conductor.beat_number = 0
conductor.bar_number = 0
conductor.has_offset_adjusted = false

-- Helper to calculate crotchet duration
function conductor.get_crotchet()
	return 60 / conductor.bpm
end

-- Update function to handle timing
function conductor.update(dt)
	-- Update song position
	conductor.song_position = conductor.song_position + dt

	-- Calculate next beat and bar times
	conductor.next_beat_time = conductor.song_position + conductor.get_crotchet()
	conductor.next_bar_time = conductor.song_position + (conductor.get_crotchet() * conductor.crotchets_per_bar)

	-- Optionally, logic for beat/bar events can go here
end

-- Snap time to the nearest beat
function conductor.snap_to_beat(time)
	local crotchet = conductor.get_crotchet()
	return math.floor((time + conductor.offset) / crotchet) * crotchet
end

function conductor.get_bpm()
	return conductor.bpm
end


return conductor
