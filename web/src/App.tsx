import React, {useState} from 'react';
import 'react-notifications/lib/notifications.css';
import axios from 'axios'
// @ts-ignore
import {NotificationContainer, NotificationManager} from 'react-notifications'
import styled from 'styled-components'

const Box = styled.div`
  background: linear-gradient(to bottom right, #eee, #fff);
  width: 100vw;
  height: 99vh;
  display: flex;
  justify-content: center;
  align-items: center;
`

const AnswerP = styled.div`
  text-align: center;
  padding: 10px;
  font-size: 14px;

  & strong {
    font-size: 18px;
  }
`

const InnerBox = styled.div`
`

const H1 = styled.h1`
  display: block;
  font-size: 16px;
  text-align: center;
`

const Input = styled.input`
  flex-basis: 100%;
  height: 30px;
  width: 300px;
  border: solid .5px gray;
  padding: 10px;
  font-size: 14px;
  border-radius: 10px;

  &:focus {
    outline: none;
  }

`

const Button = styled.button`
  display: block;
  margin: 10px auto 10px;
  border: solid .5px white;
  background-color: green;
  color: white;
  font-size: 14px;
  height: 40px;
  width: 150px;
  box-shadow: -2px 2px gray;
  cursor: pointer;
  border-radius: 2.5px;
  transition: all .5s ease-in-out;

  &:active {
    box-shadow: -3px 3px gray;

  }

  &:disabled {
    background-color: #879087;
  }
`


function App() {

    const [input, setInput] = useState("")
    const [highestPrimeNumber, setHighestPrimeNumber] = useState(0)
    const [loading, setLoading] = useState(false)

    function changeInput(e: React.ChangeEvent<HTMLInputElement>) {
        if (Number.isNaN(Number(e.target.value))) {
            return
        }

        setInput(e.target.value)
    }

    async function submit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault()
        setLoading(true)

        try {
            const {data} = await axios.post("/api", {number: Number(input)})
            setHighestPrimeNumber(data.number)
        } catch (err) {
            if (err.response) {
                NotificationManager.error(err.response.data.error)
            } else {
                NotificationManager.error("Check your network connection")
            }
            console.log(err)
        }

        setLoading(false)

    }

    return (
        <>
            <NotificationContainer/>
            <Box>
                <InnerBox>
                    <H1>Find The Highest Prime Number</H1>
                    <form onSubmit={submit}>
                        <Input
                            type={"text"}
                            placeholder={"Input a number to find the highest"}
                            required={true}
                            value={input}
                            onChange={changeInput}
                        />
                        <AnswerP>Highest Prime Number is <strong>{highestPrimeNumber || ""}</strong></AnswerP>
                        <Button type={"submit"} disabled={loading}>
                            {loading ? 'Loading...' : 'Find'}
                        </Button>
                    </form>
                </InnerBox>
            </Box>
        </>
    );
}

export default App;
