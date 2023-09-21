const localStorageMock = (() => {
    let store: Record<string, string> = {};
  
    return {
      getItem(key: string): string | null {
        return store[key] || null;
      },
  
      setItem(key: string, value: string): void {
        store[key] = value;
      },
  
      clear(): void {
        store = {};
      },
  
      removeItem(key: string): void {
        delete store[key];
      },
  
      getAll(): Record<string, string> {
        return store;
      },
    };
  })();

export const mockLocalStorage = () => {
    Object.defineProperty(window, "localStorage", { value: localStorageMock });
}

export const mockingFetch = (statuscode: number, data?: object) => {
  const mock = jest.spyOn(global, 'fetch').mockResolvedValue(
      new Response (
          JSON.stringify(data),
          {
              status: statuscode,
          }
  ));
  return mock;
}

export const mockingFetchJson = (data: object) => {
  const mock = global.fetch = jest.fn().mockResolvedValue({
    json: () => Promise.resolve(data),
  });
  return mock;
}