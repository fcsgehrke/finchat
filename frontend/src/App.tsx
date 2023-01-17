import type { Component } from 'solid-js';
import Login from './pages/Login'
import 'bootstrap/dist/css/bootstrap.min.css';
import Chats from './pages/Chats';
import Chat from './pages/Chat';
import Signup from './pages/Signup';
import Email from './pages/Email';
import Confirmed from './pages/Confirmed';
import RoomCreate from './pages/RoomCreate';

const App: Component = () => {
  return (
    <>
      <div style={'height: 100vh; display: flex; align-items: center; justify-content: center;'} class='container'>
        <RoomCreate />
      </div>
    </>
  );
};

export default App;
