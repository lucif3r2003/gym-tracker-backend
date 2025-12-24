import express from "express";
import { addExerciseController } from "../controllers/exercise_controller.js";

const router = express.Router();

router.post("/workout-day/:dayId/exercises", addExerciseController);

export default router;
