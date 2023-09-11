import { screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

export const loginFormFill = (username: string, password: string) => {
    userEvent.type(screen.getByLabelText('username'), username)
    userEvent.type(screen.getByLabelText('password'), password)
}

