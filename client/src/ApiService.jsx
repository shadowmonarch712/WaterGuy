import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:8000',
  headers: {
    'Content-Type': 'application/json'
  }
});
export const createUserEntry = (userID, value) => {
    return apiClient.post('/store', { userID, value });
  };
export const fetchUserValue = (userID) => {
  return apiClient.get(`/fetch?UserID=${userID}`);
};

export const updateUserValue = (userID, value) => {
  return apiClient.put('/update', { userID, value });
};

export const deleteUser = (userID) => {
  return apiClient.delete(`/delete?userID=${userID}`);
};
