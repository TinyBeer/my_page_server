import { createSlice } from '@reduxjs/toolkit';
import axios from 'axios';

const memoStore = createSlice({
  name: 'memo',
  initialState: {
    memoList: [],
  },
  reducers: {
    setMemoList(state, action) {
      state.memoList = action.payload;
    },
  },
});

const { setMemoList } = memoStore.actions;

function fetchMemoList(success, failed) {
  return (dispatch) => {
    const access_token = localStorage.getItem('access_token');
    const instance = axios.create({
      baseURL: 'http://127.0.0.1:9999',
      timeout: 1000,
      headers: {
        'Content-Type': 'application/json',
        Authorization: access_token,
      },
    });
    instance
      .get('/memo/list')
      .then((res) => res.data)
      .then((data) => {
        if (data.status === 'success') {
          dispatch(setMemoList(data.memoes));
          if (success) {
            success();
          }
        } else {
          throw new Error(data.message);
        }
      })
      .catch((error) => {
        console.error('list memo error:', error);
        if (failed) {
          failed();
        }
      });
  };
}

const reducer = memoStore.reducer;

export { fetchMemoList };
export default reducer;
