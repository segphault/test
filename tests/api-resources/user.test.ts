// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

import PetstoreFix from 'petstore-fix';
import { Response } from 'node-fetch';

const petstoreFix = new PetstoreFix({
  apiKey: 'My API Key',
  oauthAccessToken: 'My OAuth Access Token',
  baseURL: process.env['TEST_API_BASE_URL'] ?? 'http://127.0.0.1:4010',
});

describe('resource user', () => {
  test('create', async () => {
    const responsePromise = petstoreFix.user.create();
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('create: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.user.create({ path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('create: request options and params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.user.create(
        {
          id: 10,
          email: 'john@email.com',
          firstName: 'John',
          lastName: 'James',
          password: '12345',
          phone: '12345',
          username: 'theUser',
          userStatus: 1,
        },
        { path: '/_stainless_unknown_path' },
      ),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });

  test('deleteUsername', async () => {
    const responsePromise = petstoreFix.user.deleteUsername('string');
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('deleteUsername: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.user.deleteUsername('string', { path: '/_stainless_unknown_path' }),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });

  test('listWithCreate: only required params', async () => {
    const responsePromise = petstoreFix.user.listWithCreate([{}, {}, {}]);
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('listWithCreate: required and optional params', async () => {
    const response = await petstoreFix.user.listWithCreate([
      {
        id: 10,
        username: 'theUser',
        firstName: 'John',
        lastName: 'James',
        email: 'john@email.com',
        password: '12345',
        phone: '12345',
        userStatus: 1,
      },
      {
        id: 10,
        username: 'theUser',
        firstName: 'John',
        lastName: 'James',
        email: 'john@email.com',
        password: '12345',
        phone: '12345',
        userStatus: 1,
      },
      {
        id: 10,
        username: 'theUser',
        firstName: 'John',
        lastName: 'James',
        email: 'john@email.com',
        password: '12345',
        phone: '12345',
        userStatus: 1,
      },
    ]);
  });

  test('login', async () => {
    const responsePromise = petstoreFix.user.login();
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('login: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.user.login({ path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('login: request options and params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.user.login(
        { password: 'string', username: 'string' },
        { path: '/_stainless_unknown_path' },
      ),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });

  test('logout', async () => {
    const responsePromise = petstoreFix.user.logout();
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('logout: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.user.logout({ path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('retrieveUsername', async () => {
    const responsePromise = petstoreFix.user.retrieveUsername('string');
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('retrieveUsername: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.user.retrieveUsername('string', { path: '/_stainless_unknown_path' }),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });

  test('updateUsername: only required params', async () => {
    const responsePromise = petstoreFix.user.updateUsername({ path_username: 'string' });
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('updateUsername: required and optional params', async () => {
    const response = await petstoreFix.user.updateUsername({
      path_username: 'string',
      id: 10,
      email: 'john@email.com',
      firstName: 'John',
      lastName: 'James',
      password: '12345',
      phone: '12345',
      body_username: 'theUser',
      userStatus: 1,
    });
  });
});
