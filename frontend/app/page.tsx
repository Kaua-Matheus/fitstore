import { Default } from './assets/consumers/goApplication';

export default async function Home() {
  const response = await Default();
  console.log(response);

  return (
    <div>
      Olá mundo!
    </div>
  )
}