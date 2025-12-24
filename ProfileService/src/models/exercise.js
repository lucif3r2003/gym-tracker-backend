import mongoose from "mongoose";

const ExerciseSchema = new mongoose.Schema({
  slug: {
    type: String,
    required: true,
    index: true,
  },
  name: {
    type: String,
    required: true,
  },
  primary_muscle: {
    type: String,
    required: true,
  },
  secondary_muscles: {
    type: [String],
    default: [],
  },
  category: {
    type: String,
  },
  mechanics: {
    type: String,
  },
  force: {
    type: String,
  },
  difficulty: {
    type: String,
  },
}, {
  timestamps: true, // optional
  collection: "exercises",
});

export default mongoose.model("Exercise", ExerciseSchema);
