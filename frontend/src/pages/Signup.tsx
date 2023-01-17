
import Logo from '../components/Logo';

export default function Signup() {
    return (
            <div style={'display: flex; text-align: center; justify-content: center; align-items: center;'}>
                <div style={'width: 28rem;'}>
                    <Logo />
                    <form action="">
                        <div class='p-2 mt-4'>
                            <div class='mb-2'>
                                <label class='h5'>Basic Data</label>
                                <input class='form-control' type="text" placeholder='John Smith' required />
                            </div>
                            <div class='mb-4'>
                                <input class='form-control' type="email" placeholder='email@example.com' required />
                            </div>
                            <div class='mb-2'>
                                <label class='h5'>Password</label>
                                <input class='form-control' type="password" placeholder='Password' required />
                            </div>
                            <div class='mb-2'>
                                <input class='form-control' type="password" placeholder='Password Confirm' required />
                            </div>
                            <div class='d-grid gap-2 mt-4'>
                                <button class='btn btn-dark btn-lg'>
                                    Confirm
                                </button>
                                <button class='btn btn-white btn-lg'>
                                    Cancel
                                </button>
                            </div>
                        </div>
                    </form>
                </div>
            </div>)
}