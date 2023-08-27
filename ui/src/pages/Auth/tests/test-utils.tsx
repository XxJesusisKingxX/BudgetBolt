import { fireEvent, getByLabelText } from "@testing-library/react";

export const loginFormFill = (username: string, password: string, element: HTMLElement) => {
    const user = getByLabelText(element, "username");
    const pass = getByLabelText(element, "password");
    fireEvent.change(user, { target: { value: username }});
    fireEvent.change(pass, { target: { value: password }});
}

export const mockingFetch = (statuscode: number) => {
    const mock = jest.spyOn(global, 'fetch').mockResolvedValue(
        new Response (
            JSON.stringify({
                id: 1
            }),
            {
                status: statuscode,
            }
    ));
    return mock;
}