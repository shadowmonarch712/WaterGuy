import React from 'react'
import Form from './Components/form'
const App = () => {
  return (<>
    {/* <Reminder/> */}
    <div class="text-center my-8">
      <h1 class="text-4xl font-bold text-blue-600">Welcome to WaterGuy</h1>
      <p class="text-xl text-gray-600 mt-2">Your personal water intake reminder</p>
    </div>

    <Form />
  </>
  )
}

export default App