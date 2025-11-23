import './App.css';
import InputPage from './pages/InputPage';
import ResultsPage from './pages/ResultsPage';
import { useState, useEffect } from 'react';
import { Parse } from '../wailsjs/go/app/App';
import { types } from '../wailsjs/go/models';
import { EventsOn } from '../wailsjs/runtime/runtime';

function App() {
  const [input, setInput] = useState<string>("");
  const [parserResult, setParserResult] = useState<types.ParserResult | null>(null);
  const [ready, setReady] = useState<boolean>(false);

  useEffect(() => {
    EventsOn('parserReady', () => {
      setReady(true);
    });
  }, []);

  useEffect(() => {
    if (!input) {
      return;
    }
    const run = async () => {
      const res = await Parse(input);
      setParserResult(res);
    };
    run();
  }, [input]);

  if (ready === false) {
    return(
      <p className='text-6xl text-center mt-32'>Парсер загружается...</p>
    );
  } else if (input === "") {
    return (
      <InputPage onSubmit={(v: string) => setInput(v)} />
    );
  } else {
    return (
      <ResultsPage parserResult={parserResult} reset={() => setInput("")} />
    )
  }
}

export default App;
