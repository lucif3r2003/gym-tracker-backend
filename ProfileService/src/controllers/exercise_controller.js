import { addExerciseToDay, findExerciseById, findWorkoutDayById } from "../repository/exercise_repo.js"

export const addExerciseController = async(req, res, next) =>{
  try {
    const {dayId} = req.params;
    const exerciseData = req.body
    if (!dayId || !exerciseData){
      return res.status(400).json({
        message: "Invalid Request",
      })
    }
  // 1️⃣ check workout day 
    const day = await findWorkoutDayById(dayId);
    if (!day) {
      return res.status(404).json({
        success: false,
        message: "Workout day not found",
      });
    }

    // 2️⃣ check exercise
    const exercise = await findExerciseById(exerciseData.exercise_id);
    if (!exercise) {
      return res.status(404).json({
        success: false,
        message: "Exercise not found",
      });
    }

    // 3️⃣ add exercise
    const updatedDay = await addExerciseToDay(dayId, exerciseData);

    return res.status(200).json({
      success: true,
      data: updatedDay,
    });
    
  } catch (error) {
    next(error)  
  }
}
