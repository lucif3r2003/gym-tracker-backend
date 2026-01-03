
import mongoose from "mongoose";

const WorkoutScheduleSchema = new mongoose.Schema({
  name: {
    type: String,
    required: true, // "Push Pull Legs"
  },


  days: [
    {
      day_of_week: {
        type: Number,
        required: true, // 1 = Monday ... 7 = Sunday
      },
      workout_day_id: {
        type: mongoose.Schema.Types.ObjectId,
        ref: "WorkoutDay",
        required: true,
      },
    }
  ],

  duration_weeks: {
    type: Number,
    default: 4,
  },

  user_id: {
    type: mongoose.Schema.Types.ObjectId,
    ref: "User",
    required: true,
  },

  is_active: {
    type: Boolean,
    default: true,
  },

}, {
  timestamps: true,
  collection: "workout_schedules",
});

export default mongoose.model("WorkoutSchedule", WorkoutScheduleSchema);
