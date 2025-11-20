import { Button } from "@/components/ui/button"
import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"
import { types } from '../../wailsjs/go/models'
import { useState } from 'react';

interface ResultsPageProps {
  parserResult: types.ParserResult | null;
  reset: () => void;
}

const POSStyles = new Map<string, string>([
  ['名詞', 'text-orange-400'], //noun
  ['動詞', 'text-green-500'], //verb
  ['形容詞', 'text-blue-500'], //i ajective
  ['副詞', 'text-yellow-300'], //adverb
  ['連体詞', 'text-blue-500'], //rentaishi ???
  ['助動詞', 'text-purple-500'], //auxillary verb
  ['助詞', 'text-gray-300'], //particle
  ['接続詞', 'text-gray-300'], //conjunction
]);

function ResultsPage({ parserResult, reset } : ResultsPageProps) {
  const [showFurigana, setShowFurigana] = useState<boolean>(true);
  const [showRomaji, setShowRomaji] = useState<boolean>(false);
  const [showPolivanov, setShowPolivanov] = useState<boolean>(false);

  return(
    <div className="w-full grid grid-cols-1 justify-items-center mx-auto py-8 ">
      <div className="flex flex-wrap gap-x-2 border border-solid border-secondary text-4xl mx-8 p-4">
        {parserResult ? parserResult.Tokens.map((token, key) => {
          if ((token.Surface.match(/\n/g) || []).length != 0) {
            return <div key={key} className="basis-full h-0"></div>;
          } else return (
            <div key={key} className={`flex flex-col items-center ${POSStyles.has(token.POS[0]) ? POSStyles.get(token.POS[0]) : 'text-gray-300'}`}>
              {(token.POS[0] !== "記号" && token.Reading !== '*') && (
                <>
                  {showPolivanov && <span className='text-sm'>{token.Polivanov}</span>}
                  {showRomaji && <span className='text-sm'>{token.Romaji}</span>}
                  {token.Surface !== token.Reading && (
                    <>
                      {showFurigana && <span className='text-sm'>{token.Reading}</span>}
                    </>
                  )}
                </>
              )}
              <span className='mt-auto'>{token.Surface}</span>
            </div>
          )}
        ) : (<p>чтото пошло не так</p>)}
      </div>
      <div className='flex flex-wrap gap-4 mt-4'>
        <div className="flex items-center space-x-2">
          <Switch id="furigana" checked={showFurigana} onCheckedChange={(checked: boolean) => setShowFurigana(checked)} />
          <Label htmlFor="furigana">Показывать фуригану</Label>
        </div>
        <div className="flex items-center space-x-2">
          <Switch id="romaji" checked={showRomaji} onCheckedChange={(checked: boolean) => setShowRomaji(checked)} />
          <Label htmlFor="romaji">Показывать ромадзи</Label>
        </div>
        <div className="flex items-center space-x-2">
          <Switch id="polivanov" checked={showPolivanov} onCheckedChange={(checked: boolean) => setShowPolivanov(checked)} />
          <Label htmlFor="polivanov">Показывать киридзи</Label>
        </div>
      </div>
      <Button
        onClick={reset}
      >
        Сбросить
      </Button>
    </div>
  );
}

export default ResultsPage;
