import { useState } from 'react';
import './App.css';

function App() {
  const [name, setName] = useState('');
  const [greeting, setGreeting] = useState('');

  return (
      <div className="greeting">
        <label>Name:</label>
        <input placeholder="John Doe" onChange={(e) => {setName(e.target.value)}}/>


        <button onClick={() => {
          fetch(`http://localhost:3000/greetings/${name}`)
          .then(res => res.json())
          .then(json => {
            setGreeting(json.greeting);
          })
        }}>Submit</button>

        <label>Greeting:</label>
        <span>{greeting}</span>
      </div>
  )
}

export default App
