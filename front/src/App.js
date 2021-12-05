import logo from './media/logo.svg';
import './styles/app.css';

function App() {
  return (
    <main>
      <header>
        <img src={logo} className="app-logo" alt="logo" />
        <h1>M U D</h1>
        <a
          className="app-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </main>
  );
}

export default App;
