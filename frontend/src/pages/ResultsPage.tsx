import { Button } from "@/components/ui/button"
import { Switch } from "@/components/ui/switch"
import { Label } from "@/components/ui/label"
import { Separator } from "@/components/ui/separator"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
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
    {parserResult ? (
      <>
      <div className="flex flex-wrap gap-x-2 border border-solid border-secondary text-4xl mx-8 p-4">
        {parserResult.Tokens.map((token, key) => {
          if ((token.Surface.match(/\n/g) || []).length != 0) {
            return <div key={key} className="basis-full h-0"></div>;
          } else return (
            <div key={key} className={`flex flex-col items-center ${POSStyles.has(token.POS[0]) ? POSStyles.get(token.POS[0]) : 'text-gray-300'}`}>
              {(token.POS[0] !== "記号" && token.Reading !== '*') && (
                <>
                  {showPolivanov && <span className='text-sm'>{token.Polivanov}</span>}
                  {showRomaji && <span className='text-sm'>{token.Romaji}</span>}
                  {(showFurigana && token.Surface !== token.Reading) && (
                    <span className='text-sm'>{token.Reading}</span>
                  )}
                </>
              )}
              <Popover>
                <PopoverTrigger className='mt-auto hover:underline underline-offset-[6px] decoration-3'>{token.Surface}</PopoverTrigger>
                <PopoverContent className='w-auto max-w-4xl' align='start' collisionPadding={20}>
                  <div className='flex gap-x-4'>
                    <div className='flex flex-col place-items-center whitespace-nowrap'>
                      {token.BaseForm !== token.BaseFormReading && (
                        <p className='text-xs'>{token.BaseFormReading !== '*' ? token.BaseFormReading : token.Reading}</p>
                      )}
                      <p className='text-3xl'>{token.BaseForm}</p>
                    </div>
                    <div className='flex flex-col gap-y-2'>
                      <p className='whitespace-nowrap'>Часть речи: {token.POS[0]}{token.POS[1] !== '*' && ', ' + token.POS[1]}</p>
                      <Separator />
                      {token.Translations !== null ? token.Translations[0].Translations.map((translation, key) => (
                        <p key={key}>{translation}</p>
                      )) : (
                        <p>Значение не найдено</p>
                      )}
                      {token.InflectionalForm !== '*' && (
                        <>
                          <Separator />
                          <p>Спряжение:</p>
                          <p className='whitespace-nowrap'>{token.Surface} — {token.InflectionalForm}</p>
                        </>
                      )}
                    </div>
                  </div>
                </PopoverContent>
              </Popover>
            </div>
          )})}
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
      <div className='flex flex-wrap gap-4 mt-8 pl-8 w-fit'>
        <div className='border border-solid border-secondary text-center p-4'>
          <p className='text-2xl mb-4'>Соотношение символов</p>
          <p className='text-xl'>Хирагана: {parserResult.HKKRatio.hiragana.toFixed(2)}%</p>
          <p className='text-xl'>Катакана: {parserResult.HKKRatio.katakana.toFixed(2)}%</p>
          <p className='text-xl'>Кандзи: {parserResult.HKKRatio.kanji.toFixed(2)}%</p>
        </div>
        <div className='flex flex-col place-items-center border border-solid border-secondary p-4'>
          <p className='text-2xl mb-4'>Соотношение частей речи</p>
          <div className='flex w-full'>
            <span className='text-xl text-center w-1/2'>Часть речи</span>
            <span className='text-xl text-center w-1/2'>Подклассы</span>
          </div>
          <div className='flex flex-col gap-y-4'>
            {Object.entries(parserResult.POSStats.ExtendedRatio).map(([POS, subs]) => (
              <>
                <Separator />
                <div key={POS} className='grid grid-cols-2'>
                  <div className='flex flex-wrap place-items-center min-w-1/2 gap-x-2'>
                    <span className={`text-2xl ${POSStyles.has(POS) ? POSStyles.get(POS) : 'text-gray-300'}`}>{POS}</span>
                    <span className='text-xl'>{parserResult.POSStats.BasicRatio[POS].toFixed(2)}%</span>
                  </div>
                  <div className='flex flex-col place-items-center min-w-1/2'>
                    {Object.entries(subs).map(([subPOS, value]) => subPOS !== '*' ? (
                      <div className='flex flex-wrap place-items-center gap-x-2'>
                        <span className='text-xl'>{subPOS}</span>
                        <span className='text-lg'>{parseFloat(value as string).toFixed(2)}%</span>
                      </div>
                    ) : (
                      <span>—</span>
                    ))}
                  </div>
                </div>
              </>
            ))}
          </div>
        </div>
      </div>
    </>
    ) : (
      <p>ожидайте</p>
    )}
      <Button onClick={reset} >
        Сбросить
      </Button>
    </div>
  );
}

export default ResultsPage;
