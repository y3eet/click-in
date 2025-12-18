export type Prettify<T> = {
  [K in keyof T]: T[K];
} & {};

type Success<T> = {
  data: T;
  error: null;
};
export type ResponseError = {
  response: {
    data: {
      error: string;
    };
    status: number;
  };
};
export type Failure<E = ResponseError> = {
  data: null;
  error: E;
};

export type Result<T, E = Error> = Success<T> | Failure<E>;
