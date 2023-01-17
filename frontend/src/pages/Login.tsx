import Logo from '../components/Logo';

export default function Login() {
    return (
            <div style={'display: flex; text-align: center; justify-content: center; align-items: center;'}>
                <div style={'width: 28rem;'}>
                    <Logo />
                    <form action="">
                        <div class='p-2 mt-4'>
                            <div class='mb-2'>
                                <input class='form-control form-control-lg' type="email" placeholder='E-mail Address' required />
                            </div>
                            <div class='mb-2'>
                                <input class='form-control form-control-lg' type="password" placeholder='Password' required />
                            </div>
                            <div class='d-grid gap-2'>
                                <button class='btn btn-dark btn-lg'>
                                    Sign In
                                </button>
                                <button class='btn btn-white btn-lg'>
                                    Sign Up
                                </button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>)
}