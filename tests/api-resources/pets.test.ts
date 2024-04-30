// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

import PetstoreFix from 'petstore-fix';
import { Response } from 'node-fetch';

const petstoreFix = new PetstoreFix({
  apiKey: 'My API Key',
  oauthAccessToken: 'My OAuth Access Token',
  baseURL: process.env['TEST_API_BASE_URL'] ?? 'http://127.0.0.1:4010',
});

describe('resource pets', () => {
  test('create: only required params', async () => {
    const responsePromise = petstoreFix.pets.create({
      name: 'doggie',
      photoUrls: ['string', 'string', 'string'],
    });
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('create: required and optional params', async () => {
    const response = await petstoreFix.pets.create({
      name: 'doggie',
      photoUrls: ['string', 'string', 'string'],
      id: 10,
      category: { id: 1, name: 'Dogs' },
      status: 'available',
      tags: [
        { id: 0, name: 'string' },
        { id: 0, name: 'string' },
        { id: 0, name: 'string' },
      ],
    });
  });

  test('retrieve', async () => {
    const responsePromise = petstoreFix.pets.retrieve(0);
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
    await expect(petstoreFix.pets.retrieve(0, { path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('update: only required params', async () => {
    const responsePromise = petstoreFix.pets.update({
      name: 'doggie',
      photoUrls: ['string', 'string', 'string'],
    });
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('update: required and optional params', async () => {
    const response = await petstoreFix.pets.update({
      name: 'doggie',
      photoUrls: ['string', 'string', 'string'],
      id: 10,
      category: { id: 1, name: 'Dogs' },
      status: 'available',
      tags: [
        { id: 0, name: 'string' },
        { id: 0, name: 'string' },
        { id: 0, name: 'string' },
      ],
    });
  });

  test('delete', async () => {
    const responsePromise = petstoreFix.pets.delete(0);
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('delete: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.pets.delete(0, { path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('findByStatus', async () => {
    const responsePromise = petstoreFix.pets.findByStatus();
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('findByStatus: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.pets.findByStatus({ path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('findByStatus: request options and params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.pets.findByStatus({ status: 'available' }, { path: '/_stainless_unknown_path' }),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });

  test('findByTags', async () => {
    const responsePromise = petstoreFix.pets.findByTags();
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('findByTags: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.pets.findByTags({ path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('findByTags: request options and params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.pets.findByTags(
        { tags: ['string', 'string', 'string'] },
        { path: '/_stainless_unknown_path' },
      ),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });

  test('updateById', async () => {
    const responsePromise = petstoreFix.pets.updateById(0);
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('updateById: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.pets.updateById(0, { path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('updateById: request options and params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.pets.updateById(
        0,
        { name: 'string', status: 'string' },
        { path: '/_stainless_unknown_path' },
      ),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });

  test('uploadImage', async () => {
    const responsePromise = petstoreFix.pets.uploadImage(0);
    const rawResponse = await responsePromise.asResponse();
    expect(rawResponse).toBeInstanceOf(Response);
    const response = await responsePromise;
    expect(response).not.toBeInstanceOf(Response);
    const dataAndResponse = await responsePromise.withResponse();
    expect(dataAndResponse.data).toBe(response);
    expect(dataAndResponse.response).toBe(rawResponse);
  });

  test('uploadImage: request options instead of params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(petstoreFix.pets.uploadImage(0, { path: '/_stainless_unknown_path' })).rejects.toThrow(
      PetstoreFix.NotFoundError,
    );
  });

  test('uploadImage: request options and params are passed correctly', async () => {
    // ensure the request options are being passed correctly by passing an invalid HTTP method in order to cause an error
    await expect(
      petstoreFix.pets.uploadImage(0, { additionalMetadata: 'string' }, { path: '/_stainless_unknown_path' }),
    ).rejects.toThrow(PetstoreFix.NotFoundError);
  });
});
