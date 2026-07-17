# Calculator Frontend

## Goal

Build a production-quality calculator frontend using React and TypeScript.

The application should communicate with the backend through a REST API.

The frontend must **never evaluate mathematical expressions locally**.

Every calculation should be performed by the backend.

The application should be responsive, accessible, well-tested, and follow modern React best practices.

---

# Technology Stack

- React 19+
- TypeScript
- Vite
- Axios
- React Testing Library
- Vitest
- CSS Modules
- ESLint
- Prettier
- Docker
- GitHub Actions

---

# Backend API

POST /calculate

Request

{
    "expression":"2+2*2"
}

Response

{
    "result":6
}

Error

{
    "error":"invalid expression"
}

---

# Project Structure

frontend/

src/

components/

    Calculator/

    Display/

    Keypad/

    Button/

    History/

    ErrorMessage/

    LoadingIndicator/

hooks/

    useCalculator.ts

services/

    calculatorApi.ts

reducers/

    calculatorReducer.ts

types/

utils/

styles/

tests/

public/

Dockerfile

docker-compose.yml

vite.config.ts

eslint.config.js

README.md

.github/

    workflows/

        ci.yml

---

# Architecture

Application

↓

Calculator Page

↓

Calculator Component

↓

Reducer

↓

API Service

↓

Backend

The UI should not know anything about Axios.

All API communication must happen inside services/.

---

# Calculator Layout

The calculator should look similar to a standard calculator.

--------------------------------

Display

--------------------------------

7 8 9 /

4 5 6 *

1 2 3 -

0 . = +

C DEL ( )

--------------------------------

The user may either

- click buttons

or

- type directly into the display.

---

# Features

Support entering

numbers

decimal values

parentheses

addition

subtraction

multiplication

division

power

percentage

sqrt()

The frontend should only build the expression.

The backend performs evaluation.

---

# State Management

Use useReducer instead of multiple useState hooks.

State

{
    expression:"",
    result:null,
    loading:false,
    error:null
}

---

# Reducer Actions

SET_EXPRESSION

APPEND_CHARACTER

DELETE_CHARACTER

CLEAR

SET_LOADING

SET_RESULT

SET_ERROR

---

# API Layer

Never use Axios directly inside components.

Create

services/calculatorApi.ts

Expose

calculate(expression)

The service should

send HTTP request

return typed response

throw typed errors

---

# Custom Hook

Create

useCalculator()

Responsibilities

manage reducer

handle API calls

handle loading

handle errors

expose actions to UI

Components should remain mostly presentational.

---

# Components

Calculator

Top level component.

Coordinates child components.

---

Display

Shows

current expression

latest result

loading indicator

Should support keyboard input.

---

Keypad

Contains all calculator buttons.

No business logic.

---

Button

Reusable calculator button.

Supports

number

operator

action

variant styling

---

History

Optional.

Display the last five calculations.

History should only exist in memory.

No persistence.

---

ErrorMessage

Shows backend validation errors.

No alert().

---

LoadingIndicator

Simple spinner.

Displayed while waiting for backend.

---

# Keyboard Support

Support

0-9

+

-

*

/

.

(

)

Backspace

Delete

Escape

Enter

Keyboard interaction should behave exactly like clicking buttons.

---

# Validation

Frontend validates

empty expression

expression length

invalid repeated operators where obvious

Backend remains source of truth.

---

# Loading State

While request is pending

disable = button

show spinner

prevent duplicate requests

---

# Error Handling

Handle

network failures

backend validation

timeouts

unexpected server errors

Display friendly messages.

---

# Accessibility

Target WCAG 2.2 AA compliance.

---

## Semantic HTML

Use

<main>

<form>

<button>

<output>

Avoid unnecessary divs.

---

## Labels

Interactive controls require accessible labels.

Buttons containing only symbols should include

aria-label

Examples

"Add"

"Subtract"

"Equals"

"Delete"

"Clear"

---

## Screen Reader Support

Result

role="status"

aria-live="polite"

Errors

role="alert"

aria-live="assertive"

Screen readers should automatically announce

new results

validation errors

network errors

---

## Keyboard Navigation

Tab order must be logical.

Every button should be reachable.

Enter triggers calculation.

Space activates focused buttons.

Visible focus indicators required.

---

## Focus Management

After calculation

keep focus on the display.

On validation errors

move focus to the display.

---

## Touch Targets

Buttons

minimum

44px

---

## Contrast

Minimum contrast ratio

4.5:1

Do not rely solely on color.

---

## Reduced Motion

Respect

prefers-reduced-motion

---

# Responsive Design

Support

Desktop

Tablet

Mobile

Calculator should remain centered.

Buttons should resize gracefully.

No horizontal scrolling.

---

# Styling

Use CSS Modules.

Modern clean design.

Rounded corners.

Subtle shadows.

Responsive spacing.

Consistent typography.

No UI component libraries.

---

# Testing

Use

Vitest

React Testing Library

Write tests for

Calculator component

Reducer

API service

Custom hook

Keyboard interaction

Loading state

Error rendering

Accessibility

Button clicks

Display rendering

Target

90%+ coverage.

---

# Docker

Create

Dockerfile

Requirements

Node LTS

Production build

Serve static files

Minimal image

Also create

docker-compose.yml

for local development.

---

# Linting

Configure

ESLint

Prettier

TypeScript strict mode

No warnings.

---

# GitHub Actions

Create CI workflow.

Run on every push and pull request.

Steps

Install dependencies

Type check

Lint

Run tests

Generate coverage

Build production bundle

Fail if any step fails.

---

# README

Generate

Overview

Architecture

Folder Structure

Setup

Running Locally

Docker

Testing

Accessibility

Design Decisions

Future Improvements

Screenshots placeholder

API Integration

---

# Design Decisions

The frontend should intentionally remain thin.

It should

collect user input

render UI

call backend

display results

It must not contain calculation logic.

Business logic belongs entirely to the backend.

---

# Future Extensibility

The architecture should make it easy to add

Dark mode

Calculation history persistence

Scientific calculator

Themes

Internationalization

Offline mode

PWA support

Without significant refactoring.

---

# Code Quality

Requirements

Strict TypeScript

Functional components only

Reusable components

Custom hooks

No duplicated logic

No inline styles

Small focused files

Meaningful naming

Composition over large components

---

# IMPORTANT

Before generating code

1. Produce an implementation plan.

2. List every file to be created.

3. Explain the responsibility of each file.

4. Generate code incrementally.

5. Ensure each file compiles before moving to the next.

6. Generate production-ready code.

7. Do not leave TODOs or placeholders.

8. Ensure the application passes linting and tests.