defmodule RnaTranscription do
  @doc """
  Transcribes a character list representing DNA nucleotides to RNA

  ## Examples

  iex> RnaTranscription.to_rna('ACTG')
  'UGAC'
  """
  @spec to_rna([char]) :: [char]
  def to_rna(dna) do
    transcribe = %{
      ?A => ?U,
      ?C => ?G,
      ?T => ?A,
      ?G => ?C
    }

    dna
    |> Enum.map(fn x -> transcribe[x] end)
  end
end
