// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

import PetstoreFix from 'petstore-fix';
import { Response } from 'node-fetch';

const petstoreFix = new PetstoreFix({
  apiKey: 'My API Key',
  oauthAccessToken: 'My OAuth Access Token',
  baseURL: process.env['TEST_API_BASE_URL'] ?? 'http://127.0.0.1:4010',
});

describe('resource order', () => {
  test('retrieve', async () => {
    const responsePromise = petstoreFix.store.order.retrieve(0);
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('retrieve: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.store.order.retrieve(0, { path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('deleteOrder', async () => {
    const responsePromise = petstoreFix.store.order.deleteOrder(0);
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('deleteOrder: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.store.order.deleteOrder(0, { path: '/_stainless_unknown_path' }),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });
});
