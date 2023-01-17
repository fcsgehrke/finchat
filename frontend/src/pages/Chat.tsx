import Logo from "../components/Logo"
import Message from "../components/Message"

export default function Chat() {
    return (
        <div style={'width: 90%; align-items: center; height: 80vh;'}>
            <div class="row">
                <div class="col text-center">
                    <Logo />
                </div>
            </div>
            <div class="row text-center mt-4">
                <div class="col-3 pb-0 pt-3 text-center">
                    <p class="display-6">People</p>   
                </div>
                <div class="col ps-0 pb-0 pt-3">
                    <p class="display-6">Messages</p>   
                </div>
                <div class="col-2 float-end pe-0 mt-2">
                    <button class="btn btn-white float-end mb-2 mt-4">Exit Chat</button>
                </div>
            </div>
            <div class="row">
                <div class="col">
                    <ul class="list-group">
                        <li class="list-group-item">Person 1</li>
                        <li class="list-group-item">Person 1</li>
                        <li class="list-group-item">Person 1</li>
                        <li class="list-group-item">Person 1</li>
                        <li class="list-group-item">Person 1</li>
                    </ul>
                </div>
                <div class="col-9">
                    <div class='row' style={'height: 36em;'}>
                        <div class="col-12" style="overflow: auto; height: 36em;">
                            <Message name="Felipe C. Gehrke" time={new Date()} msg="Lorem" />
                            <Message type='mine' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                            <Message type='mine' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                            <Message type='mine' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                            <Message type='mine' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                            <Message type='mine' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                            <Message type='mine' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                            <Message type='mine' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                            <Message type='bot' name="Felipe C. Gehrke" time={new Date()} msg="Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum." />
                        </div>
                    </div>
                    <div class="row ps-0 pt-2 pb-2 mt-1" style={'height: 3.5em;'}>
                        <div class="col ps-2 pe-0" style={'display: flex;'}>
                            <input type="text" class="form-control" />
                            <button type="button" class="ms-2 btn btn-dark">Send</button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}
