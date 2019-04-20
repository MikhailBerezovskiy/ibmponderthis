/*
 * Solution for the puzzle:
 *
 * Find nine different prime numbers that can be placed
 * in a 3x3 square in such a way that the average of every
 * row, column, and diagonal is also a prime number.
 *
 * using a depth first search with limited depth
 * depth limitation is hard coded as maximum number of primes in main func
 *
 * game struct is representing a state of the game, for example:
 * initial state is
 *
 * Board:
 * [0,0,0]
 * [0,0,0]
 * [0,0,0]
 * Index: -1 current index in multidimentional array(board), can be from 0 to 8
 *      index 1 is a board[0][1], index 4 is board[1][1]
 *
 * AvMoves: available moves for next state, is a vector of the prime numbers
 * usedMoves: hashmap of primes that have been used already on the board
 *    is used for finding duplicates
 *
 * For each state algorithm is filtering next available moves for next index
 * and applying depth first search (LIFO)
 *
 */

#include <array>
#include <fstream>
#include <iostream>
#include <map>
#include <stack>
#include <vector>

using namespace std;

using board = array<array<int, 3>, 3>;

bool isPrime(int p) {
  if (p == 1) {
    return true;
  }
  for (int i = 2; i < p; ++i) {
    if (p % i == 0) {
      return false;
    }
  }
  return true;
}

// generate list of primes by given number of primes
vector<int> makePrimes(int lim) {
  vector<int> vp;
  int primeNum = 0;
  for (int i = 0; i < lim; ++i) {
    ++primeNum;
    while (!isPrime(primeNum)) {
      ++primeNum;
    }
    vp.push_back(primeNum);
  }
  return vp;
}

// takes 3 integers and check if average is prime
bool avgIsPrime(int a, int b, int c) {
  int sum = a + b + c;
  if (sum % 3 == 0 && isPrime(sum / 3)) {
    return true;
  }
  return false;
}

// game object, represent a single state of the game
struct game {
  bool win;                 // win!
  int ind;                  // current index on board 0..8
  board b;                  // board [3][3]int
  vector<int> avMoves;      // available moves, or next primes
  map<int, bool> usedMoves; // primes that are assigned

  game() {
    for (int i = 0; i < 3; ++i) {
      b[i] = array<int, 3>{0, 0, 0};
    }
    win = false;
    ind = -1;
    avMoves = vector<int>{};
    usedMoves = map<int, bool>{};
  }

  game(board oldB, vector<int> am, map<int, bool> um) {
    win = false;
    b = oldB;
    avMoves = am;
    usedMoves = um;
  }

  // assign new value(prime number) to current index
  void setMove(int move) {
    // assign move to the current index
    b[ind / 3][ind % 3] = move;
  }

  // filter is taking current available moves and coordinates(indexes) of
  // another 2 elements of row(column, or diagonal)
  // returns new list of available moves
  //
  // for example: avmoves=[5,7,9]. testing row with [1,3,...]
  // only 5 is valid number to satisfy initial game instructions
  // so function will return [5]
  //
  vector<int> filter(vector<int> &mv, int x1, int y1, int x2, int y2) {
    vector<int> nmv;
    for (auto i : mv) {
      if (isDup(i)) {
        continue;
      }
      if (avgIsPrime(b[x1][y1], b[x2][y2], i)) {
        nmv.push_back(i);
      }
    }
    return nmv;
  }

  // check if prime number already used on the board
  bool isDup(int val) {
    if (usedMoves.find(val) == usedMoves.end()) {
      return false;
    }
    return true;
  }

  // depends on current index, check constraints for the next move
  // it will significantly reduce search space
  void setAvMoves(vector<int> oldm) {
    avMoves = oldm;
    switch (ind) {
    case 1:
      // row 1
      avMoves = filter(avMoves, 0, 0, 0, 1);
      break;
    case 4:
      // row 2
      avMoves = filter(avMoves, 1, 0, 1, 1);
      break;
    case 5:
      // col 1
      avMoves = filter(avMoves, 0, 0, 1, 0);
      // diag '/'
      avMoves = filter(avMoves, 0, 2, 1, 1);
      break;
    case 6:
      // col 2
      avMoves = filter(avMoves, 0, 1, 1, 1);
      break;
    case 7:
      // diag '\'
      avMoves = filter(avMoves, 0, 0, 1, 1);
      // row 3
      avMoves = filter(avMoves, 2, 0, 2, 1);
      // col 3
      avMoves = filter(avMoves, 0, 2, 1, 2);
      break;

    default:
      vector<int> newMoves;
      for (auto i : oldm) {
        if (isDup(i)) {
          continue;
        }
        newMoves.push_back(i);
      }
      avMoves = newMoves;
      break;
    }
  }

  void printFunc() {
    cout << "board: ind=" << ind << endl;
    for (int i = 0; i < 3; ++i) {
      cout << b[i][0] << '\t' << b[i][1] << '\t' << b[i][2] << endl;
    }
    cout << "aval moves size=" << avMoves.size() << endl;
  }
};

// create new game state by copying previous state
// and new prime number for the previous state
game newGame(game &g, int move) {

  // construct a new state
  game ng = game(g.b, g.avMoves, g.usedMoves);

  // assign new prime to next the index on the board
  ng.ind = g.ind + 1;
  ng.setMove(move);
  ng.usedMoves.insert({move, true});

  // filter next moves
  ng.setAvMoves(g.avMoves);

  return ng;
}

// play is a recursive function with depth first search algorithm
game play(game &g) {

  // WIN condition:
  // propagete win backward
  if (g.win) {
    return g;
  }
  // WIN found 9 numbers (0..8)
  if (g.ind == 8) {
    g.win = true;
    return g;
  }

  // Lose condition
  // no available moves
  if (g.avMoves.empty()) {
    g.win = false;
    return g;
  }

  for (auto i : g.avMoves) {
    // for each new available move(prime number), create a child state
    game ng = newGame(g, i);
    game result = play(ng);
    if (result.win) {
      return result;
    }
  }

  // lost
  return g;
}

//
//
int main(void) {
  ofstream resFile;
  resFile.open("solutions.txt");

  // initialize list of primes, that will be used for finding solution
  vector<int> primes = makePrimes(40);
  for (auto i : primes) {
    cout << i << " ";
  }
  cout << endl;

  // go through all primes and start game for each one
  for (auto p : primes) {
    // constructor for the first game state
    game gg = game();
    gg.avMoves = primes;

    game g = newGame(gg, p);

    game result = play(g);
    if (result.win) {
      cout << "WIN" << endl;
      result.printFunc();

      for (int i = 0; i < 3; ++i) {
        resFile << result.b[i][0] << '\t' << result.b[i][1] << '\t'
                << result.b[i][2] << endl;
      }
      resFile << endl;
    }
  }
  resFile.close();
  return 0;
}
