#include <string>
#include <algorithm>
#include <fstream>
#include <iostream>
#include <optional>
#include <tuple>
#include <vector>

std::optional<std::tuple<int, int>> sum_two( const std::vector<int>& ints, int sum ) {
    for ( int front = 0, back = ints.size()-1; front < back; ) {
        if ( ints[front] + ints[back] == sum ) return std::tuple<int,int>{ ints[front], ints[back] };
        if ( ints[front] + ints[back] < sum ) front++;
        else back--;
    }
    return std::nullopt;
}

std::optional<std::tuple<int,int,int>> sum_three( const std::vector<int>& ints, int sum ) {
    for ( auto item : ints ) {
        auto result = sum_two( ints, sum - item );
        if ( result.has_value() ) {
            auto [ fst, snd ] = result.value();
            return std::tuple<int, int, int>{ fst, snd, item };
        }
    }
    return std::nullopt;
}

int main() {
    std::fstream file;
    file.open( "./../inputs/1.txt", std::ios::in );

    std::vector<int> ints;
    for ( int item = 0; file >> item; ) ints.emplace_back( item );
    std::sort( ints.begin(), ints.end() );

    auto result_sum_two = sum_two( ints, 2020 );
    if ( result_sum_two.has_value() ) {
        auto [ fst, snd ] = result_sum_two.value();
        std::cout << "sum_two: " << std::to_string(fst * snd) << std::endl;
    } else {
        std::cout << "did not find sum of two for 2020" << std::endl;
    }

    auto result_sum_three = sum_three( ints, 2020 );
    if ( result_sum_three.has_value() ) {
        auto [ fst, snd, thd ] = result_sum_three.value();
        std::cout << "sum_three: " << std::to_string(fst * snd * thd) << std::endl;
    } else {
        std::cout << "did not find sum of three for 2020" << std::endl;
    }
}