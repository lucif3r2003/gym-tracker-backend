
import mongoose from "mongoose";

const WorkoutDaySchema = new mongoose.Schema({
  name: {
    type: String,
    required: true, // Ví dụ: "Chest + Triceps"
  },

  description: String,

  exercises: [
    {
      exercise_id: {
        type: mongoose.Schema.Types.ObjectId,
        ref: "Exercise",
        required: true,
      },
      sets: { type: Number, required: true },
      reps: { type: Number, required: true },
      rest_seconds: { type: Number, default: 90 },
      order: { type: Number, required: true },
      note: String,
    }
  ],

  created_by: {
    type: mongoose.Schema.Types.ObjectId,
    ref: "User",
    default: null, // null = template mặc định
  },

}, {
  timestamps: true,
  collection: "workout_days",
});

export default mongoose.model("WorkoutDay", WorkoutDaySchema);
