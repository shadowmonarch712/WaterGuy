import React, { useState, useEffect, useContext } from 'react';
import { GoalContext } from '../Context/GoalContext';
import WaterAudio from '../assets/waterReminder.mp3'

const WaterIntakeReminder = ({ setShowForm }) => {
  const [showStartPrompt, setShowStartPrompt] = useState(true);
  const [showSetLimit, setShowSetLimit] = useState(false);
  const [limit, setLimit] = useState(1);
  const { goal, setGoal } = useContext(GoalContext);
  const [timer, setTimer] = useState(null);
  const [message, setMessage] = useState('');

  const handleStartYes = () => {
    if (goal > 0) {
      setShowForm(false);
      setShowStartPrompt(false);
      setShowSetLimit(true);
    } else {
      setMessage('Please set a goal first.');
    }
  };

  const handleStartNo = () => {
    setMessage('Let me know when you want to start.'); 
  };

  const handleSetLimit = (e) => {
    setLimit(e.target.value);
  };

  const startTimer = () => {
    setShowSetLimit(false);
    if (timer) clearInterval(timer);

    const newTimer = setInterval(() => {
        setGoal((prevGoal) => {
          if (prevGoal > 1) {
            return prevGoal - 1;
          } else if (prevGoal === 1) {
            clearInterval(newTimer);
            playAudio();
            return 0; // Goal reached
          }
        });
      }, limit * 3600000); // Convert hours to milliseconds

    setTimer(newTimer);
  };

  const playAudio = () => {
    const audio = new Audio(WaterAudio);
    audio.play();
  };

  useEffect(() => {
    // Cleanup timer on component unmount
    return () => {
      if (timer) clearInterval(timer);
    };
  }, [timer]);

  return (
    <div className="container mx-auto p-4 text-center">
      {showStartPrompt && (
        <div className="my-4">
          <p className="text-lg mb-4">Do you want to start?</p>
          <button
            className="bg-green-500 text-white px-4 py-2 rounded-md hover:bg-green-700 mr-2"
            onClick={handleStartYes}
          >
            Yes
          </button>
          <button
            className="bg-red-500 text-white px-4 py-2 rounded-md hover:bg-red-700"
            onClick={handleStartNo}
          >
            No
          </button>
          {message && <p className="text-lg mt-4">{message}</p>} {/* Display message */}
        </div>
      )}

      {showSetLimit && (
        <div className="my-4">
          <p className="text-lg mb-4">Set limit in hours:</p>
          <input
            className="border-2 border-gray-300 p-2 rounded-md mb-4"
            type="number"
            value={limit}
            onChange={handleSetLimit}
          />
          <button
            className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-700"
            onClick={startTimer}
          >
            Start Timer
          </button>
        </div>
      )}

{!showStartPrompt && !showSetLimit && goal === 0 && (
        <div className="my-4">
          <p className="text-lg mb-4 font-bold">Congratulations for completing your goal!</p>
        </div>
      )}
      {!showStartPrompt && !showSetLimit && goal > 0 && (
        <div className="my-4">
          <p className="text-lg mb-4">{goal} glasses left</p>
        </div>
      )}
    </div>
  );
};

export default WaterIntakeReminder;
