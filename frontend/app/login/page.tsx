'use client'

import { useFormState, useFormStatus } from 'react-dom';
import { authenticate } from '@/app/lib/actions';

const initialState = {
    username: '',
    password: ''
}

export default function Login(): JSX.Element {

    const [_, dispatch] = useFormState(authenticate, initialState);

    return (
       <div>
            <h2 className="text-2xl flex font-bold place-content-center">Login</h2>
            <form className="flex flex-col items-center" action={dispatch} method="POST">
               <input className="p-1 m-1 w-40 h-8 rounded text-black" name="username" type="text" placeholder="Username" required /> 
               <input className="p-1 m-1 w-40 h-8 rounded text-black" name="password" type="password" placeholder="Password" required /> 
               <LoginButton />
            </form>

            {/* Left in to make the background color of the page fill the entire screen */}
            <div className="h-screen">
            
            </div>

       </div>
    )
}

function LoginButton(): JSX.Element {

    const { pending } = useFormStatus();
    
    const handleClick = (event: any) => {
        if (pending) {
            event.preventDefault();
        }
    }

    return (
            <button 
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-2"
                type = "submit"
                onClick={handleClick}
                aria-disabled={pending}
            >
                Login
            </button>
    );
}
