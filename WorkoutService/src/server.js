import express from 'express'
import dotenv from 'dotenv'
import ExcerciseRoute from "./routes/ExerciseRoute.js"

dotenv.config()

const app = express();
const PORT = process.env.PORT

app.use(express.json())
app.use('/api/workout', ExcerciseRoute)
app.listen(PORT, ()=>{
  console.log(`Profile service start with port : ${PORT}`)
})
