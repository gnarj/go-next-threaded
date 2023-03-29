/** @type {import('ts-jest').JestConfigWithTsJest} */
module.exports = {
  preset: 'ts-jest',
  testEnvironment: 'jsdom',
  moduleNameMapper: {
    // '\\.(css|less)$': '<rootDir>/test/jest/__mocks__/styleMock.js',
    '\\.(css|less)$': '<rootDir>/styleMock.js',
  },
  setupFilesAfterEnv: ['<rootDir>/setupTests.ts'],
};
