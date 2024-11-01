import dogsService from '../service/dogs'
import {useState, useEffect} from 'react'
import Media from './media'
import MediaMock from './mediaMock'
import utils from '../utils'
import { useHistory } from "react-router-dom";

function Home(){
    const [dog, setDog] = useState('')
    const [loading, setLoading] = useState(false)
    const [count, setCount] = useState(utils.getDogs().length)

    const history = useHistory();

    // Keep updated the dogs hook
    useEffect(()=>{
        setLoading(true)
        dogsService.getDog().then(r => {
            setDog(r)
            setLoading(false)
        })
    }, {dog})

    // Redirect if dogs are selected
    const redirect = ()=>{
        if (count < 5) return

        if (utils.getDogs().length === 0){
            alert('Hi! Try to choose a dog please!')
            setCount(0)
        }

        if (utils.getDogs().length !== 0){
            history.push('favorites')
        }
    }

    // Load the next dog
    const onNextClicked = ()=>{
        redirect()

        setLoading(true)

        dogsService.getDog().then(r => {
            setCount(count + 1)
            setDog(r)
            setLoading(false)
        })
    }

    // Save the dog and load the next dog
    const onSaveClicked = ()=> {
        utils.saveDog(dog)

        onNextClicked()
    }

    return (<div className="home shadow">
        <h1>Select the dog #{count + 1} 🎉</h1>
        {loading ? <MediaMock/> : <Media dog={dog}/>}
        <button style={{'justify-self' : 'flex-end'}} onClick={()=>onSaveClicked()}>
            Save it 🤩
        </button>
        <button style={{'justify-self' : 'flex-start'}} onClick={()=>onNextClicked()}>
            Next 🙆
        </button>
    </div>)
}

export default Home