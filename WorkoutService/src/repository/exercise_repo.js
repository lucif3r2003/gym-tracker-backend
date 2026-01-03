import workout_day from "../models/workout_day.js";
import exercise from "../models/exercise.js";

//logic handle -----------------------------------------------------------------------------
//find exercise by id
export const findExerciseById = (id) =>{
  return exercise.findById(id) 
}

//find workOut day by id
export const findWorkoutDayById = (id) => {
  return workout_day.findById(id)
}

//workout handle ---------------------------------------------------------------------------
//add exercise to a day
export const addExerciseToDay = async (dayId, exerciseData) => {
  try {
   return workout_day.findOneAndUpdate(
      dayId,
      {
        $push:{
          exercises: exerciseData,
        },
      },
    ) 
  } catch (error) {
    logger.error("Failed to add exercise to workout day", {
      dayId,
      exerciseData,
      error: error.message,
      stack: error.stack,
    });

    throw error;
  } 
}

//add a new workout day
export const newWorkoutDay = (workoutData) =>{
  try {
   return workout_day.insertOne(workoutData)
  } catch (error) {
    
  }
}
