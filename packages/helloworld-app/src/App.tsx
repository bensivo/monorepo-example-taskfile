import { useState } from 'react';
import * as Proto from 'helloworld-proto-ts';
import './App.css';

function App() {
  const [name, setName] = useState('');
  const [greeting, setGreeting] = useState('');

  return (
    <div className="greeting">
      <label>Name:</label>
      <input placeholder="John Doe" onChange={(e) => { setName(e.target.value) }} />

      <button onClick={() => {

        const greetRequest = new Proto.GreetRequest({
          name: name,
        });

        fetch(`http://localhost:3000/greet`, {
          method: "POST",
          body: greetRequest.toBinary(),
        })
          .then(res => res.arrayBuffer())
          .then(arrBuffer => Proto.GreetResponse.fromBinary(new Uint8Array(arrBuffer)))
          .then(greetResponse => {
            setGreeting(greetResponse.greeting);
          })
      }}>Submit</button>

      <label>Greeting:</label>
      <span>{greeting}</span>
    </div>
  )
}

export default App
