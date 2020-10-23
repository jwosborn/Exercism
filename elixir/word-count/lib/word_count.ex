defmodule WordCount do
  @doc """
  Count the number of words in the sentence.

  Words are compared case-insensitively.
  """
  @spec count(String.t()) :: map
  def count(sentence) do
    sentence
    |> to_string
    |> String.downcase()
    |> String.replace(~r/[^\w\d\s_-]/u, "")
    |> String.split(
      ~r/_|\s/,
    trim: true
    )
    |> Enum.reduce(%{}, &counter/2)
  end

  def counter(str, acc) do
    Map.get(acc, str)
    |> has(str, acc)
  end

  def has(nil, key, acc) do
    Map.put(acc, key, 1)
  end

  def has(value, key, acc) do
    Map.put(acc, key, value + 1)
  end
end
